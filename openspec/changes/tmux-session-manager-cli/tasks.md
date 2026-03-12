## 1. Project Bootstrap

- [ ] 1.1 Initialize Go module and Cobra command structure for the `tm` CLI entrypoint
- [ ] 1.2 Add required dependencies (Cobra, Bubble Tea/Bubbles list components, TOML parser) and verify build compiles
- [ ] 1.3 Add repository README notes for local usage and planned tmux workflow behavior

## 2. Configuration System

- [ ] 2.1 Implement config path resolution using XDG config home with fallback to `~/.config/tm/config.toml`
- [ ] 2.2 Implement TOML config model with `search_path` field and load/save helpers
- [ ] 2.3 Implement create-if-missing behavior for `.config/tm/` directory and `config.toml`
- [ ] 2.4 Add tests for first-run config creation and config update behavior

## 3. Config CLI Commands

- [ ] 3.1 Implement `tm config` parent command and nested `set` subcommand in Cobra
- [ ] 3.2 Implement `tm config set search-path <path>` argument parsing and path validation
- [ ] 3.3 Ensure invalid paths return clear errors without overwriting existing valid config
- [ ] 3.4 Add command tests covering success, invalid path, and file-creation-on-first-write flows

## 4. Repository Discovery and Selection

- [ ] 4.1 Implement repository discovery scanning immediate child directories for `.git` metadata
- [ ] 4.2 Implement root command default flow to load config and discover repositories
- [ ] 4.3 Implement the repository selector UI using the Bubble Tea `list-simple` example as the interaction model
- [ ] 4.4 Render repo list items with repo name title and full path description
- [ ] 4.5 Implement non-TTY behavior with clear fallback/error messaging
- [ ] 4.6 Add tests for discovery filtering, selector filtering, selection success path, and cancel/non-interactive handling

## 5. Worktree Selection and Tmux Launch

- [ ] 5.1 Implement Git worktree inspection for a selected repository using Git metadata
- [ ] 5.2 Show a second Bubble Tea selector with main checkout plus linked worktrees when multiple targets exist
- [ ] 5.3 Skip the second selector when only the main checkout is available
- [ ] 5.4 Implement deterministic tmux session naming based on selected repo/worktree path
- [ ] 5.5 Implement attach-or-create tmux behavior rooted at the selected target path
- [ ] 5.6 Add tests for worktree choice generation, session naming, and tmux error handling

## 6. Git and Delivery Setup

- [ ] 6.1 Initialize or verify local Git repository is on `main` branch
- [ ] 6.2 Create initial commit(s) for scaffold and implemented features
- [ ] 6.3 Add GitHub remote and push branch to origin when credentials and repository are configured
