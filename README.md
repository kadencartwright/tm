# tm

`tm` is a small Go CLI for jumping from a repository picker into the right tmux session.

## Requirements

- **fzf** - The fuzzy finder. Install via:
  - macOS: `brew install fzf`
  - Ubuntu/Debian: `sudo apt install fzf`
  - Arch: `sudo pacman -S fzf`
  - Other: see https://github.com/junegunn/fzf#installation

## Local usage

1. Install dependencies with `go mod tidy`.
2. Run tests with `go test ./...`.
3. Start the CLI with `go run .`.
4. Set the repository search path first with `go run . config set search-path ~/code`.

## Repository Selector

When you run `tm`, you'll see an fzf-powered repo selector:

- **Type** to filter repos immediately
- **↑/↓** arrow keys to navigate the filtered list
- **Enter** to select and open the repo
- **Esc** or **Ctrl+C** to cancel

You can customize fzf behavior by setting `FZF_DEFAULT_OPTS` environment variable.

## Nested Session Handling

`tm` can be run from inside an existing tmux session. When you do this, it will:
1. Automatically detach you from the current tmux session
2. Create or attach to the requested session
3. Seamlessly transition you to the new session

This prevents the "sessions should be nested with care" error that tmux normally shows.

## Planned workflow

- `tm` loads `search_path` from `~/.config/tm/config.toml` or `$XDG_CONFIG_HOME/tm/config.toml`.
- It scans the immediate child directories under that path and lists repositories that contain `.git` metadata.
- It shows an fzf selector with repository name and full path.
- If the selected repository has linked worktrees, it shows a second selector for the checkout target.
- It attaches to an existing tmux session for that target or creates one rooted at the selected path.
