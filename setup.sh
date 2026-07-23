#!/usr/bin/env bash
set -euo pipefail

DOTCLAUDE_DIR="$(cd "$(dirname "$0")" && pwd)"
TARGET_DIR="$HOME/.claude"
BIN_DIR="$TARGET_DIR/bin"

FILES=(
  CLAUDE.md
  settings.json
  discord-notify.sh
)

echo "dotclaude setup"
echo "  source: $DOTCLAUDE_DIR"
echo "  target: $TARGET_DIR"
echo ""

# Build statusline binary
echo "Building statusline..."
if command -v go &>/dev/null; then
  mkdir -p "$DOTCLAUDE_DIR/bin"
  go build -o "$DOTCLAUDE_DIR/bin/statusline" "$DOTCLAUDE_DIR/cmd/statusline/"
  echo "  OK   bin/statusline built"
else
  echo "  WARN go not found, skipping statusline build"
fi
echo ""

mkdir -p "$TARGET_DIR" "$BIN_DIR"

# Link files
for file in "${FILES[@]}"; do
  src="$DOTCLAUDE_DIR/$file"
  dst="$TARGET_DIR/$file"

  if [[ ! -f "$src" ]]; then
    echo "  SKIP $file (source not found)"
    continue
  fi

  if [[ -L "$dst" ]]; then
    current_target="$(readlink "$dst")"
    if [[ "$current_target" == "$src" ]]; then
      echo "  OK   $file (already linked)"
      continue
    else
      echo "  WARN $file is symlinked to: $current_target"
      read -rp "       Overwrite? [y/N] " ans
      if [[ "$ans" != [yY] ]]; then
        echo "       Skipped."
        continue
      fi
      rm "$dst"
    fi
  elif [[ -f "$dst" ]]; then
    echo "  WARN $file already exists (regular file)"
    if diff -q "$src" "$dst" >/dev/null 2>&1; then
      echo "       Content is identical."
      read -rp "       Replace with symlink? [y/N] " ans
    else
      echo "       Content differs:"
      diff --color=auto -u "$dst" "$src" | head -20 || true
      echo ""
      read -rp "       Replace with symlink? (original backed up to $dst.bak) [y/N] " ans
      if [[ "$ans" == [yY] ]]; then
        cp "$dst" "$dst.bak"
      fi
    fi
    if [[ "$ans" != [yY] ]]; then
      echo "       Skipped."
      continue
    fi
    rm "$dst"
  fi

  ln -s "$src" "$dst"
  echo "  LINK $file -> $src"
done

# Link statusline binary
if [[ -f "$DOTCLAUDE_DIR/bin/statusline" ]]; then
  src="$DOTCLAUDE_DIR/bin/statusline"
  dst="$BIN_DIR/statusline"

  if [[ -L "$dst" ]] && [[ "$(readlink "$dst")" == "$src" ]]; then
    echo "  OK   bin/statusline (already linked)"
  elif [[ -e "$dst" ]]; then
    echo "  WARN bin/statusline already exists"
    read -rp "       Overwrite? [y/N] " ans
    if [[ "$ans" == [yY] ]]; then
      rm -f "$dst"
      ln -s "$src" "$dst"
      echo "  LINK bin/statusline -> $src"
    else
      echo "       Skipped."
    fi
  else
    ln -s "$src" "$dst"
    echo "  LINK bin/statusline -> $src"
  fi
fi

# Link skills
SKILLS_DIR="$TARGET_DIR/skills"
mkdir -p "$SKILLS_DIR"
for skill in "$DOTCLAUDE_DIR/skills/"*.md; do
  [[ -f "$skill" ]] || continue
  name="$(basename "$skill")"
  dst="$SKILLS_DIR/$name"

  if [[ -L "$dst" ]] && [[ "$(readlink "$dst")" == "$skill" ]]; then
    echo "  OK   skills/$name (already linked)"
  elif [[ -e "$dst" ]]; then
    echo "  WARN skills/$name already exists"
    read -rp "       Overwrite? [y/N] " ans
    if [[ "$ans" == [yY] ]]; then
      rm -f "$dst"
      ln -s "$skill" "$dst"
      echo "  LINK skills/$name -> $skill"
    else
      echo "       Skipped."
    fi
  else
    ln -s "$skill" "$dst"
    echo "  LINK skills/$name -> $skill"
  fi
done

echo ""
echo "Done."
echo ""
echo "NOTE: discord-notify.sh requires \$CLAUDE_CODE_DISCORD_WEBHOOK_URL to be set."
echo "      Add it to your shell profile (e.g. ~/.config/fish/conf.d/secrets.fish)"
