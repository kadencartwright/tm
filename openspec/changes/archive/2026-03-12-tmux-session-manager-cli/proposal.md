## Why

Developers who manage many projects lose time manually attaching to or creating tmux sessions per repository, and there is no single CLI to discover repos and jump directly into session workflows. Building a focused `tm` CLI now enables faster context switching and standardizes project navigation with a configurable repo search path.

## What Changes

- Add a new Go CLI application named `tm` built with Cobra.
- Add a Bubble Tea list-based default command flow modeled after the `list-simple` example, where invoking `tm` opens a fuzzy selector of repositories discovered from configured search paths.
- Add a second selection step for repositories that have Git worktrees, allowing the user to choose between the main checkout and linked worktrees before opening.
- Add tmux session attach/create behavior rooted at the selected repository or worktree path.
- Add configuration commands under `config`, including `config set search-path <path>`.
- Add TOML-based persistent configuration stored under `.config/tm/`, including creation of the config file and parent directories when missing.
- Add bootstrap documentation/tasks to initialize a Git repository on `main` and prepare first push to GitHub.

## Capabilities

### New Capabilities
- `repo-session-selection`: Discover repositories from configured paths and present a Bubble Tea fuzzy selection flow when `tm` is invoked.
- `worktree-selection`: Detect Git worktrees for a selected repository and present a second selection flow when multiple checkout targets exist.
- `tmux-target-opening`: Attach to or create a tmux session for the selected repository or worktree path.
- `config-search-path-management`: Manage search path configuration through CLI commands and persistent TOML config storage.

### Modified Capabilities


## Impact

- New codebase structure for a Go CLI (`cmd/`, internal packages, config handling, repo discovery, fuzzy UI integration).
- Integration with Bubble Tea list-based terminal UI and Git worktree inspection commands.
- New runtime dependencies for Cobra, TOML parsing/writing, and fuzzy selection terminal UI.
- Local developer environment impact through creation of `.config/tm/config.toml`.
- Developer workflow impact through documented steps to initialize Git on `main` and connect/push to a GitHub remote.
