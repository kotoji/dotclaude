package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"os"
	"os/exec"
	"strings"
	"time"
)

const (
	colorReset     = "\033[0m"
	colorBold      = "\033[1m"
	colorCyan      = "\033[36m"
	colorBrightBlue = "\033[94m"
	colorGreen     = "\033[32m"
	colorYellow    = "\033[33m"
	colorRed       = "\033[31m"
	colorDim       = "\033[2m"

	iconFolder = ""            // nf-md-folder (U+20A83 in Nerd Font)
	iconBranch = "\xee\x9c\xa5" // nf-dev-git_branch (U+E725)
)

type Input struct {
	Model         *ModelInfo     `json:"model"`
	ContextWindow *ContextWindow `json:"context_window"`
	Cost          *Cost          `json:"cost"`
	RateLimits    *RateLimits    `json:"rate_limits"`
	Workspace     *Workspace     `json:"workspace"`
	Effort        *Effort        `json:"effort"`
	CWD           string         `json:"cwd"`
	Version       string         `json:"version"`
}

type ModelInfo struct {
	DisplayName string `json:"display_name"`
	ID          string `json:"id"`
}

type ContextWindow struct {
	UsedPercentage    float64 `json:"used_percentage"`
	ContextWindowSize int64   `json:"context_window_size"`
	TotalInputTokens  int64   `json:"total_input_tokens"`
	TotalOutputTokens int64   `json:"total_output_tokens"`
}

type Cost struct {
	TotalCostUSD      float64 `json:"total_cost_usd"`
	TotalDurationMs   int64   `json:"total_duration_ms"`
	TotalLinesAdded   int64   `json:"total_lines_added"`
	TotalLinesRemoved int64   `json:"total_lines_removed"`
}

type RateLimits struct {
	FiveHour *RateLimit `json:"five_hour"`
	SevenDay *RateLimit `json:"seven_day"`
}

type RateLimit struct {
	UsedPercentage float64        `json:"used_percentage"`
	ResetsAt       json.RawMessage `json:"resets_at"`
}

func (rl *RateLimit) ResetTime() (time.Time, bool) {
	if len(rl.ResetsAt) == 0 {
		return time.Time{}, false
	}
	var s string
	if json.Unmarshal(rl.ResetsAt, &s) == nil && s != "" {
		t, err := time.Parse(time.RFC3339, s)
		return t, err == nil
	}
	var n float64
	if json.Unmarshal(rl.ResetsAt, &n) == nil && n > 0 {
		sec := int64(n)
		if sec > 1e12 {
			sec /= 1000
		}
		return time.Unix(sec, 0), true
	}
	return time.Time{}, false
}

type Workspace struct {
	CurrentDir string `json:"current_dir"`
}

type Effort struct {
	Level string `json:"level"`
}

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprint(os.Stdout, "statusline error: "+err.Error())
		return
	}

	var input Input
	if err := json.Unmarshal(data, &input); err != nil {
		fmt.Fprint(os.Stdout, "statusline error: "+err.Error())
		return
	}

	sep := colorDim + " | " + colorReset

	line1 := strings.Join([]string{
		shortDir(input),
		model(input),
		effort(input),
		contextBar(input),
	}, sep)

	line2 := strings.Join([]string{
		repoBranch(input),
		rateLimit5h(input),
		rateLimit7d(input),
		cost(input),
		duration(input),
	}, sep)

	fmt.Fprintf(os.Stdout, "%s\n%s", line1, line2)
}

func shortDir(input Input) string {
	cwd := input.CWD
	if input.Workspace != nil && input.Workspace.CurrentDir != "" {
		cwd = input.Workspace.CurrentDir
	}
	if cwd == "" {
		return "N/A"
	}
	parts := strings.Split(cwd, "/")
	if len(parts) == 0 {
		return colorBrightBlue + iconFolder + " " + cwd + colorReset
	}
	return colorBrightBlue + iconFolder + " …/" + parts[len(parts)-1] + colorReset
}

func model(input Input) string {
	name := "N/A"
	if input.Model != nil {
		if input.Model.DisplayName != "" {
			name = input.Model.DisplayName
		} else if input.Model.ID != "" {
			name = input.Model.ID
		}
	}
	return colorCyan + name + colorReset
}

func repoBranch(input Input) string {
	cwd := input.CWD
	if input.Workspace != nil && input.Workspace.CurrentDir != "" {
		cwd = input.Workspace.CurrentDir
	}

	repo := ""
	if cwd != "" {
		cmd := exec.Command("git", "rev-parse", "--show-toplevel")
		cmd.Dir = cwd
		if out, err := cmd.Output(); err == nil {
			parts := strings.Split(strings.TrimSpace(string(out)), "/")
			repo = parts[len(parts)-1]
		}
	}

	branchName := ""
	if cwd != "" {
		cmd := exec.Command("git", "branch", "--show-current")
		cmd.Dir = cwd
		if out, err := cmd.Output(); err == nil {
			branchName = strings.TrimSpace(string(out))
		}
	}

	if repo == "" && branchName == "" {
		return colorGreen + " < not found" + colorReset
	}
	if branchName == "" {
		return repo
	}
	return colorGreen + " " + repo + " [" + iconBranch + " " + branchName + "]" + colorReset
}

func effort(input Input) string {
	if input.Effort == nil || input.Effort.Level == "" {
		return ""
	}
	level := input.Effort.Level
	var color string
	switch level {
	case "low":
		color = colorDim
	case "medium":
		color = colorGreen
	case "high":
		color = colorYellow
	case "xhigh", "max":
		color = colorRed
	default:
		color = colorReset
	}
	return color + "effort:" + level + colorReset
}

func thresholdColor(pct float64) string {
	if pct >= 90 {
		return colorRed
	}
	if pct >= 70 {
		return colorYellow
	}
	return colorGreen
}

func blockBar(pct float64, width int) string {
	filled := int(math.Round(pct * float64(width) / 100))
	if filled > width {
		filled = width
	}
	if filled < 0 {
		filled = 0
	}
	color := thresholdColor(pct)
	return color + strings.Repeat("█", filled) + colorDim + strings.Repeat("░", width-filled) + colorReset
}

func contextBar(input Input) string {
	if input.ContextWindow == nil {
		return blockBar(0, 10) + " 0%"
	}
	pct := input.ContextWindow.UsedPercentage
	size := input.ContextWindow.ContextWindowSize
	used := int64(math.Round(float64(size) * pct / 100))

	color := thresholdColor(pct)
	return fmt.Sprintf("%s %s%d%%%s %s/%s",
		blockBar(pct, 10),
		color, int(math.Round(pct)), colorReset,
		fmtK(used), fmtK(size))
}

func rateLimit5h(input Input) string {
	if input.RateLimits == nil || input.RateLimits.FiveHour == nil {
		return "5h:N/A"
	}
	return formatRateLimit("5h", input.RateLimits.FiveHour)
}

func rateLimit7d(input Input) string {
	if input.RateLimits == nil || input.RateLimits.SevenDay == nil {
		return "7d:N/A"
	}
	return formatRateLimit("7d", input.RateLimits.SevenDay)
}

func formatRateLimit(label string, rl *RateLimit) string {
	pct := rl.UsedPercentage
	color := thresholdColor(pct)
	bar := blockBar(pct, 8)
	reset := ""
	if t, ok := rl.ResetTime(); ok {
		reset = fmtDuration(time.Until(t))
	}
	if reset != "" {
		return fmt.Sprintf("%s:%s %s%.1f%%%s (%s)", label, bar, color, pct, colorReset, reset)
	}
	return fmt.Sprintf("%s:%s %s%.1f%%%s", label, bar, color, pct, colorReset)
}

func tokens(input Input) string {
	if input.ContextWindow == nil {
		return "0 tok"
	}
	total := input.ContextWindow.TotalInputTokens + input.ContextWindow.TotalOutputTokens
	return fmtK(total) + " tok"
}

func cost(input Input) string {
	if input.Cost == nil || input.Cost.TotalCostUSD == 0 {
		return colorYellow + "$0.00" + colorReset
	}
	return colorYellow + fmt.Sprintf("$%.2f", input.Cost.TotalCostUSD) + colorReset
}

func duration(input Input) string {
	if input.Cost == nil || input.Cost.TotalDurationMs == 0 {
		return "0m"
	}
	sec := input.Cost.TotalDurationMs / 1000
	h := sec / 3600
	m := (sec % 3600) / 60
	if h > 0 {
		return fmt.Sprintf("%dh%02dm", h, m)
	}
	return fmt.Sprintf("%dm", m)
}

func fmtK(n int64) string {
	if n >= 1_000_000 {
		return fmt.Sprintf("%.1fM", float64(n)/1_000_000)
	}
	if n >= 1_000 {
		return fmt.Sprintf("%dk", n/1000)
	}
	return fmt.Sprintf("%d", n)
}

func fmtDuration(d time.Duration) string {
	if d < 0 {
		d = 0
	}
	mins := int(d.Minutes())
	if mins < 60 {
		return fmt.Sprintf("%dm", mins)
	}
	h := mins / 60
	m := mins % 60
	return fmt.Sprintf("%dh%02dm", h, m)
}
