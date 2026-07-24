#!/usr/bin/env bash
set -euo pipefail

if [[ -z "${CLAUDE_CODE_DISCORD_WEBHOOK_URL:-}" ]]; then
	echo "CLAUDE_CODE_DISCORD_WEBHOOK_URL is not set" >&2
	exit 1
fi

MESSAGE=""
MIN_DURATION=0

while [[ $# -gt 0 ]]; do
	case "$1" in
	--min-duration)
		MIN_DURATION="$2"
		shift 2
		;;
	*)
		MESSAGE="$1"
		shift
		;;
	esac
done

# Skip if task duration is below threshold
if [[ "$MIN_DURATION" -gt 0 ]] && [[ -n "${CLAUDE_PID:-}" ]]; then
	BUSY_SINCE="$(jq -r '.statusUpdatedAt // empty' "$HOME/.claude/sessions/${CLAUDE_PID}.json" 2>/dev/null || echo "")"
	if [[ -n "$BUSY_SINCE" ]]; then
		NOW_MS="$(date +%s)000"
		ELAPSED_SEC=$(((NOW_MS - BUSY_SINCE) / 1000))
		if [[ "$ELAPSED_SEC" -lt "$MIN_DURATION" ]]; then
			exit 0
		fi
	fi
fi

HOSTNAME_SHORT="$(hostname -s)"

# Repository name or current directory name
if git rev-parse --show-toplevel &>/dev/null; then
	REPO_NAME=":git: $(basename "$(git rev-parse --show-toplevel)")"
else
	REPO_NAME=":folder: $(basename "$(pwd)")/"
fi

# Git branch
BRANCH="$(git branch --show-current 2>/dev/null || echo "")"

# Session name (from Claude Code's session file)
SESSION_NAME=""
if [[ -n "${CLAUDE_PID:-}" ]]; then
	SESSION_NAME="$(jq -r '.name // empty' "$HOME/.claude/sessions/${CLAUDE_PID}.json" 2>/dev/null || echo "")"
fi

# tmux session and window
TMUX_SESSION=""
TMUX_WINDOW=""
if [[ -n "${TMUX:-}" ]]; then
	TMUX_SESSION="$(tmux display-message -p '#{session_name}' 2>/dev/null || echo "")"
	TMUX_WINDOW="$(tmux display-message -p '#{window_name}' 2>/dev/null || echo "")"
fi

# Color based on message content
if [[ "$MESSAGE" == *"Done"* ]] || [[ "$MESSAGE" == *"done"* ]] || [[ "$MESSAGE" == *"✨"* ]]; then
	COLOR=5763719 # green (0x57F287)
elif [[ "$MESSAGE" == *"error"* ]] || [[ "$MESSAGE" == *"fail"* ]] || [[ "$MESSAGE" == *"❌"* ]]; then
	COLOR=15548997 # red (0xED4245)
else
	COLOR=5793266 # blurple (0x5865F2)
fi

# Shorten home prefix to ~
DISPLAY_DIR="${PWD/#$HOME/\~}"

TIMESTAMP="$(date -u +%Y-%m-%dT%H:%M:%S.000Z)"

FIELDS='[
	{"name": "Host", "value": $host, "inline": true},
	{"name": "Repository", "value": $repo, "inline": true},
	{"name": "Branch", "value": $branch, "inline": true}
]'
if [[ -n "$TMUX_SESSION" ]]; then
	FIELDS="$FIELDS + [{\"name\": \"tmux\", \"value\": \$tmux_info, \"inline\": true}]"
fi
if [[ -n "$SESSION_NAME" ]]; then
	FIELDS="$FIELDS + [{\"name\": \"Session\", \"value\": \$session, \"inline\": true}]"
fi
#FIELDS="$FIELDS + [{\"name\": \"Directory\", \"value\": \$dir, \"inline\": false}]"

PAYLOAD=$(
	jq -n \
		--arg content "$MESSAGE [${REPO_NAME}]" \
		--arg username "claude-code@${HOSTNAME_SHORT}" \
		--arg avatar "https://raw.githubusercontent.com/lobehub/lobe-icons/refs/heads/master/packages/static-png/dark/claudecode-color.png" \
		--arg repo "$REPO_NAME" \
		--arg host "$HOSTNAME_SHORT" \
		--arg branch "$BRANCH" \
		--arg session "$SESSION_NAME" \
		--arg tmux_info "${TMUX_SESSION}${TMUX_WINDOW:+:${TMUX_WINDOW}}" \
		--arg dir "$DISPLAY_DIR" \
		--argjson color "$COLOR" \
		--arg ts "$TIMESTAMP" \
		"{
        username: \$username,
        avatar_url: \$avatar,
        content: \$content,
        embeds: [{
            color: \$color,
			      fields: ($FIELDS),
            timestamp: \$ts
			  }]
    }"
)

curl -s -H "Content-Type: application/json" \
	-d "$PAYLOAD" \
	"$CLAUDE_CODE_DISCORD_WEBHOOK_URL" &
