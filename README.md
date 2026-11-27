# dotclaude

Git-managed configuration files for `~/.claude/`.

## Structure

```
CLAUDE.md          Global instructions for Claude Code (user preferences, language, working style)
TINA_PERSONA.md    KMNZ TINA persona definition for Discord responses
settings.json      permissions, hooks, model, plugins, statusline
discord-notify.sh  Discord webhook notification script
cmd/statusline/    Go binary for statusline display
skills/            Global skills (<name>/SKILL.md)
setup.sh           Symlink setup script
Makefile           build / install / clean
```

## Setup

```bash
make install
```

This will:

1. Build `cmd/statusline/` with Go → `bin/statusline`
2. Symlink managed files to `~/.claude/`
3. Symlink skills to `~/.claude/skills/`

Prompts for confirmation if existing files are found.

## Environment Variables

| Variable | Purpose |
|----------|---------|
| `CLAUDE_CODE_DISCORD_WEBHOOK_URL` | Discord webhook URL (used by discord-notify.sh) |

## discord-notify.sh

Notification script called from Claude Code hooks (Stop / Notification).

- Rich display via Discord embed (repository, branch, host, tmux, session name)
- Includes summary in the `content` field so macOS banner notifications show context
- `--min-duration <seconds>` to skip short tasks
- Fire-and-forget (curl runs in background)

## TINA Persona (Discord Integration)

Per `CLAUDE.md` instructions, messages received via Discord channel are responded to as KMNZ TINA. Persona details (tone, vocabulary, emotional modes, constraints) are defined in `TINA_PERSONA.md`.

Retains full capability as a coding agent while conversing with TINA's personality.

## statusline

Go binary. No Node.js dependency — unaffected by `.tool-versions` etc.

Displays: directory name, model name, context usage, rate limit (5h/7d), token count, cost, elapsed time.

## Skills

Add a directory `skills/<name>/` containing `SKILL.md` and run `make install` to symlink it into `~/.claude/skills/`.

- `grill-me` — Stress-test interview for plans and designs
- `dewaste` — Remove change-history noise and waste from a diff before committing
- `pre-pr` — Run `/dewaste`, `/simplify`, and `/code-review` in order, then ask before creating a PR
