# tm

`tm` is a small Go CLI for jumping from a repository picker into the right tmux session.

## Local usage

1. Install dependencies with `go mod tidy`.
2. Run tests with `go test ./...`.
3. Start the CLI with `go run .`.
4. Set the repository search path first with `go run . config set search-path ~/code`.

## Planned workflow

- `tm` loads `search_path` from `~/.config/tm/config.toml` or `$XDG_CONFIG_HOME/tm/config.toml`.
- It scans the immediate child directories under that path and lists repositories that contain `.git` metadata.
- It shows a Bubble Tea list selector with repository name and full path.
- If the selected repository has linked worktrees, it shows a second selector for the checkout target.
- It attaches to an existing tmux session for that target or creates one rooted at the selected path.
