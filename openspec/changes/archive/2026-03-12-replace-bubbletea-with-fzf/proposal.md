## Why

The current Bubble Tea-based selector provides a custom TUI experience, but fzf is a widely-adopted standard for fuzzy finding that users already know and prefer. Replacing Bubble Tea with fzf provides a simpler, more familiar interface that integrates better with Unix tooling and requires less custom UI code to maintain.

## What Changes

- **Remove Bubble Tea dependency**: Remove `github.com/charmbracelet/bubbletea` and `github.com/charmbracelet/bubbles` from the project
- **Add fzf integration**: Implement selection using fzf as an external binary or via a Go fzf library
- **Refactor selector package**: Replace `BubbleSelector` with an `FzfSelector` implementation
- **Maintain interface compatibility**: Keep the existing `selector.Choice` struct and `Select` method signature
- **Preserve TTY detection**: Continue using `selector.IsTTY()` for conditional selection behavior
- **BREAKING**: Users must have fzf installed on their system (add to documentation)

## Capabilities

### New Capabilities
<!-- No new capabilities introduced - this is a refactoring of existing functionality -->

### Modified Capabilities

The following capabilities' implementation will change (requirements remain the same, but the underlying selector mechanism changes):

- `fuzzy-find-typing`: The fuzzy find and type selection mechanism will use fzf instead of Bubble Tea's built-in filtering. The interface and behavior remain the same (filter as you type, select with Enter).
- `repo-session-selection`: Repository selection will use fzf for the interactive picker. The flow remains the same (list repos, fuzzy find, select one).
- `worktree-selection`: Worktree target selection will use fzf. Selection behavior and filtering logic remain unchanged.

## Impact

- **Dependencies**: Remove `github.com/charmbracelet/bubbletea` and `github.com/charmbracelet/bubbles`, add dependency on fzf binary being available in PATH (or optionally a Go fzf library)
- **Binary size**: Likely reduction in binary size due to removing Bubble Tea framework
- **External dependency**: Users will need fzf installed (common on most developer systems)
- **Code changes**: `internal/selector/selector.go` complete rewrite, `cmd/root.go` may need minor adjustments if interface changes
- **Testing**: Selector tests in `cmd/root_test.go` use a fake selector, so they should continue to work with interface compatibility maintained
