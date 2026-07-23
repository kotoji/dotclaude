package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
	"time"
)

type Input struct {
	Model         *ModelInfo     `json:"model"`
	ContextWindow *ContextWindow `json:"context_window"`
	Cost          *Cost          `json:"cost"`
	RateLimits    *RateLimits    `json:"rate_limits"`
	Workspace     *Workspace     `json:"workspace"`
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
	TotalCostUSD       float64 `json:"total_cost_usd"`
	TotalDurationMs    int64   `json:"total_duration_ms"`
	TotalLinesAdded    int64   `json:"total_lines_added"`
	TotalLinesRemoved  int64   `json:"total_lines_removed"`
}

type RateLimits struct {
	FiveHour *RateLimit `json:"five_hour"`
	SevenDay *RateLimit `json:"seven_day"`
}

type RateLimit struct {
	UsedPercentage float64 `json:"used_percentage"`
	ResetsAt       string  `json:"resets_at"`
}

type Workspace struct {
	CurrentDir string `json:"current_dir"`
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

	parts := []string{
		shortDir(input),
		model(input),
		contextBar(input),
		rateLimit5h(input),
		rateLimit7d(input),
		tokens(input),
		cost(input),
		duration(input),
	}

	fmt.Fprint(os.Stdout, strings.Join(parts, " | "))
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
		return cwd
	}
	return parts[len(parts)-1]
}

func model(input Input) string {
	if input.Model != nil {
		if input.Model.DisplayName != "" {
			return input.Model.DisplayName
		}
		if input.Model.ID != "" {
			return input.Model.ID
		}
	}
	return "N/A"
}

func contextBar(input Input) string {
	if input.ContextWindow == nil {
		return "[                    ] 0% 0/0"
	}
	pct := int(math.Round(input.ContextWindow.UsedPercentage))
	size := input.ContextWindow.ContextWindowSize
	used := int64(math.Round(float64(size) * float64(pct) / 100))

	filled := pct * 20 / 100
	bar := "[" + strings.Repeat("=", filled) + strings.Repeat(" ", 20-filled) + "]"
	return fmt.Sprintf("%s %d%% %s/%s", bar, pct, fmtK(used), fmtK(size))
}

func rateLimit5h(input Input) string {
	if input.RateLimits == nil || input.RateLimits.FiveHour == nil {
		return "5h:N/A"
	}
	rl := input.RateLimits.FiveHour
	pct := rl.UsedPercentage
	bar := makeBar(pct, 15)
	reset := fmtReset(rl.ResetsAt)
	if reset != "" {
		return fmt.Sprintf("5h:%s %.1f%% (%s)", bar, pct, reset)
	}
	return fmt.Sprintf("5h:%s %.1f%%", bar, pct)
}

func rateLimit7d(input Input) string {
	if input.RateLimits == nil || input.RateLimits.SevenDay == nil {
		return "7d:N/A"
	}
	rl := input.RateLimits.SevenDay
	pct := rl.UsedPercentage
	bar := makeBar(pct, 15)
	reset := fmtReset(rl.ResetsAt)
	if reset != "" {
		return fmt.Sprintf("7d:%s %.1f%% (%s)", bar, pct, reset)
	}
	return fmt.Sprintf("7d:%s %.1f%%", bar, pct)
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
		return "$0.0000"
	}
	return fmt.Sprintf("$%.4f", input.Cost.TotalCostUSD)
}

func duration(input Input) string {
	if input.Cost == nil || input.Cost.TotalDurationMs == 0 {
		return "N/A"
	}
	sec := input.Cost.TotalDurationMs / 1000
	h := sec / 3600
	m := (sec % 3600) / 60
	return fmt.Sprintf("%dh%02dm", h, m)
}

func makeBar(pct float64, width int) string {
	filled := int(pct) * width / 100
	return "[" + strings.Repeat("=", filled) + strings.Repeat(" ", width-filled) + "]"
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

func fmtReset(iso string) string {
	if iso == "" {
		return ""
	}
	t, err := time.Parse(time.RFC3339, iso)
	if err != nil {
		return ""
	}
	diff := time.Until(t)
	if diff < 0 {
		diff = 0
	}
	mins := int(diff.Minutes())
	if mins < 60 {
		return fmt.Sprintf("%dm", mins)
	}
	h := mins / 60
	m := mins % 60
	return fmt.Sprintf("%dh%02dm", h, m)
}
