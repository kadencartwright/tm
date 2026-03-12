## Context

The project currently uses Bubble Tea (charmbracelet/bubbletea) with the bubbles list component for interactive selection. The `internal/selector` package provides a `BubbleSelector` that implements a TUI with:
- Filter-as-you-type search
- Arrow key navigation (vim keys disabled)
- Visual list with title and description
- Cancel (Esc/q/Ctrl+C) and select (Enter) actions

The selector is used in two main flows:
1. **Repository selection**: When multiple repos match, user selects from a list
2. **Target selection**: When opening a repo, user may select from multiple worktrees/targets

The current implementation is ~145 lines in `internal/selector/selector.go` with the `BubbleSelector` struct and `Select` method.

## Goals / Non-Goals

**Goals:**
- Replace Bubble Tea-based selection with fzf-based selection
- Maintain the same public interface (`Choice` struct, `Select` method signature)
- Preserve all existing functionality (filtering, selection, cancellation)
- Support TTY detection for conditional interactive mode
- Ensure tests continue to work with minimal changes

**Non-Goals:**
- Changing the selection UX significantly (keep similar feel)
- Supporting non-fzf fallback (assume fzf is available or error)
- Adding new selection features beyond what's currently supported
- Changing how choices are formatted or displayed

## Decisions

### Use External fzf Binary vs Go Library

**Decision**: Use external fzf binary via exec.Command

**Rationale**:
- fzf is already installed on most developer systems
- Using the binary provides the exact fzf experience users expect
- Avoids adding heavy Go dependencies (fzf libraries are large)
- Easier to maintain - no need to wrap fzf's complex internals
- Users can customize fzf via FZF_DEFAULT_OPTS environment variable

**Alternative considered**: Use a Go fzf library like `github.com/ktr0731/go-fuzzyfinder`
- Rejected: Adds ~1MB+ to binary, less configurable, different UX from standard fzf

### Interface Compatibility Strategy

**Decision**: Keep exact same `Selector` interface, only implementation changes

**Rationale**:
- `cmd/root.go` and tests use interface, not concrete type
- Zero changes needed in calling code
- Tests using `fakeSelector` continue to work unchanged
- Easy to swap back if needed

### Input/Output Handling

**Decision**: Pass choices to fzf via stdin, parse selection from stdout

**Rationale**:
- Standard fzf interface
- Supports filtering immediately
- Clean separation - fzf handles all UI

**Format**: JSON lines or tab-separated values
- `Label\tDetails\tValue` - user sees "Label: Details", we return Value
- Using tabs allows colons in data without collision

### Error Handling

**Decision**: Distinguish between "user cancelled" and "no match selected"

**Rationale**:
- fzf exit codes: 0 = selection made, 1 = no match, 130 = interrupted (Ctrl-C/Esc)
- Map exit code 130 to "cancelled" (same as current behavior)
- Map exit code 1 to "no selection" but not cancelled
- Map exit code 0 to successful selection

## Risks / Trade-offs

**[Risk] fzf not installed** → **Mitigation**: Check for fzf at startup, return clear error message suggesting installation. Document fzf as requirement in README.

**[Risk] Different UX feel** → **Mitigation**: fzf provides similar fuzzy filtering. Users may actually prefer it. Test with actual usage scenarios.

**[Risk] fzf version differences** → **Mitigation**: Use basic fzf features available in all versions. Standard flags (--height, --reverse) are stable.

**[Risk] Binary size increase** → **Trade-off**: Actually decreases - removing Bubble Tea saves ~500KB-1MB, fzf is external so no binary impact.

**[Risk] Windows compatibility** → **Mitigation**: fzf works on Windows. Verify in CI if needed. Current Bubble Tea also works on Windows so parity maintained.

**[Risk] Environment variable pollution** → **Mitigation**: Respect FZF_DEFAULT_OPTS but also set explicit flags to ensure consistent behavior (e.g., --height=50% --reverse).
