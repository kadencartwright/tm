## Context

The change introduces a new Go CLI that helps developers jump into tmux session workflows from a repository picker instead of manually navigating directories and session names. The CLI must support a default interactive flow (`tm`) and persistent user configuration through TOML under `.config/tm/`. The selector UX should closely follow Bubble Tea's `list-simple` example while extending it to handle a second picker for Git worktrees before opening a tmux session.

## Goals / Non-Goals

**Goals:**
- Provide a Cobra-based CLI binary named `tm`.
- Load and persist a TOML config at `.config/tm/config.toml`.
- Implement `tm config set search-path <path>` and create config directories/files when missing.
- Discover repositories under configured search paths and display them in a Bubble Tea list selector when `tm` is invoked with no subcommand.
- Detect Git worktrees for a selected repository and show a second list when more than one checkout target exists.
- Attach to an existing tmux session or create a new one rooted at the selected repository or worktree path.

**Non-Goals:**
- Multi-user/shared config support.
- Remote repository discovery (network APIs).
- Complex profile management beyond the single `search-path` key.

## Decisions

1. CLI architecture uses Cobra root command plus nested `config set` command.
   - Rationale: Cobra provides robust command parsing, help generation, and a clear extension path for future tmux/session commands.
   - Alternative considered: Custom flag parsing with `flag` package. Rejected because nested subcommands and discoverability would be weaker.

2. Config path is resolved via XDG config home with fallback to `$HOME/.config/tm/config.toml`.
   - Rationale: Matches Linux/macOS conventions while keeping behavior deterministic.
   - Alternative considered: Store config in working directory. Rejected because behavior would vary by invocation location.

3. Config schema starts with a single field: `search_path` (string path).
   - Rationale: Matches immediate requirement and keeps migration simple.
   - Alternative considered: Array of paths initially. Rejected to avoid unnecessary complexity before clear user need.

4. `config set search-path` performs idempotent create-or-update and validates path existence.
   - Rationale: Prevents broken defaults and ensures first-run success by creating missing config file and parent directory.
   - Alternative considered: Allow non-existent paths for future use. Rejected to avoid silent runtime errors in repo discovery.

5. Repository discovery scans immediate child directories and identifies repositories by presence of `.git` directory.
   - Rationale: Fast and predictable behavior aligned with common local development layouts.
   - Alternative considered: Recursive deep scan. Rejected for performance and noise concerns.

6. Interactive selection uses Bubble Tea's list model patterned after the upstream `list-simple` example.
   - Rationale: Matches the requested interaction style, reduces custom UI work, and keeps the terminal experience idiomatic.
   - Alternative considered: Another fuzzy picker library. Rejected because the Bubble Tea example gives a clearer starting point for the desired UX.

7. After repository selection, the CLI inspects `git worktree list --porcelain` and opens a second selector containing the main checkout plus linked worktrees when multiple targets exist.
   - Rationale: Lets users choose the exact checkout context before tmux attach/create while keeping the common no-worktree case fast.
   - Alternative considered: Always open the primary checkout. Rejected because it hides useful worktree context.

8. Session launch uses selected path-derived session naming and performs attach-or-create behavior.
   - Rationale: Aligns with the intended outcome of `tm` as a session manager, not only a selector.
   - Alternative considered: Print the selected path only. Rejected because it does not fulfill the tmux management goal.

## Risks / Trade-offs

- [Path validation may reject intended future directories] -> Mitigation: Document requirement that path exists and revisit if deferred creation is needed.
- [Different OS path semantics] -> Mitigation: Normalize with `filepath.Clean`, use OS-native separators, and include cross-platform tests.
- [Large directory trees can slow startup] -> Mitigation: Limit scan depth to immediate children and ignore hidden/system folders by default.
- [Fuzzy UI dependency can complicate CI tests] -> Mitigation: Isolate selection behind interface and unit test matcher/input mapping with mocks.
- [Worktree names may be ambiguous] -> Mitigation: Show title plus full path in the list item description.
- [Tmux session naming collisions across similarly named repos/worktrees] -> Mitigation: Define deterministic path-based sanitization and test collision handling.

## Migration Plan

1. Initialize repository on `main` and create initial project skeleton for Go module and Cobra commands.
2. Add configuration loading/saving package with TOML schema and file creation behavior.
3. Add `config set search-path` command and validation.
4. Add repository discovery and Bubble Tea repo selection flow.
5. Add Git worktree inspection and second-stage selection flow.
6. Add tmux attach/create behavior for the selected target path.
7. Add tests for config creation, command behavior, repository/worktree discovery, and session naming.
8. Prepare GitHub remote and perform first push once local checks pass.

Rollback strategy:
- Revert the introducing commit(s) if behavior is unstable.
- Remove generated config file manually if needed (`~/.config/tm/config.toml`) during local rollback testing.

## Open Questions

None.
