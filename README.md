# dotclaude

`~/.claude/` の設定ファイルを Git 管理するリポジトリ。

## 構成

```
CLAUDE.md          Claude Code グローバル指示
settings.json      permissions, hooks, model, plugins, statusline
discord-notify.sh  Discord webhook 通知スクリプト
cmd/statusline/    statusline 表示用 Go バイナリ
skills/            グローバル skills (.md)
setup.sh           symlink セットアップスクリプト
Makefile           build / install / clean
```

## セットアップ

```bash
make install
```

以下が行われる:

1. `cmd/statusline/` を Go でビルド → `bin/statusline`
2. 管理対象ファイルを `~/.claude/` に symlink
3. skills を `~/.claude/skills/` に symlink

既存ファイルがある場合は上書き確認が入る。

## 環境変数

| 変数名 | 用途 |
|--------|------|
| `CLAUDE_CODE_DISCORD_WEBHOOK_URL` | Discord webhook URL（discord-notify.sh が参照） |

## discord-notify.sh

Claude Code の hook (Stop / Notification) から呼ばれる通知スクリプト。

- Discord embed でリッチ表示（リポジトリ、ブランチ、ホスト、tmux、セッション名）
- `content` フィールドにサマリーを含むので macOS バナー通知でも内容がわかる
- `--min-duration <秒>` で短い処理をスキップ可能
- fire-and-forget（curl をバックグラウンド実行）

## statusline

Go 製バイナリ。node 依存を排除し `.tool-versions` 等の影響を受けない。

表示内容: ディレクトリ名、モデル名、context 使用率、rate limit (5h/7d)、トークン数、コスト、経過時間。

## Skills

`skills/` に Markdown を追加して `make install` すれば `~/.claude/skills/` にリンクされる。

- `grill-me.md` — プランや設計を stress-test するインタビュー
