## 1. Dependency Cleanup

- [x] 1.1 Remove bubbletea and bubbles from go.mod
- [x] 1.2 Run `go mod tidy` to clean up unused dependencies
- [x] 1.3 Verify no other files import bubbletea packages

## 2. FzfSelector Implementation

- [x] 2.1 Create `FzfSelector` struct with stdin/stdout configuration
- [x] 2.2 Implement fzf binary detection (`which fzf` equivalent)
- [x] 2.3 Implement choice formatting (tab-separated values: Label\tDetails\tValue)
- [x] 2.4 Implement fzf execution with exec.Command
- [x] 2.5 Handle fzf exit codes (0=success, 1=no match, 130=cancelled)
- [x] 2.6 Parse fzf output to return selected Choice
- [x] 2.7 Set fzf options (--height=50%, --reverse, --border)
- [x] 2.8 Pass choices via stdin to fzf
- [x] 2.9 Return clear error if fzf not installed

## 3. Interface Compatibility

- [x] 3.1 Keep `Choice` struct unchanged (Label, Details, Value, Title(), Description(), FilterValue())
- [x] 3.2 Keep `Select(title string, items []Choice) (Choice, bool, error)` signature
- [x] 3.3 Keep `RepoChoices()` and `TargetChoices()` helper functions
- [x] 3.4 Keep `IsTTY()` function for TTY detection
- [x] 3.5 Rename `NewBubbleSelector` to `NewFzfSelector` or keep for compatibility

## 4. Integration Updates

- [x] 4.1 Update `cmd/root.go` to use `NewFzfSelector` instead of `NewBubbleSelector`
- [x] 4.2 Verify `cmd/root_test.go` fake selector still works (should use interface)
- [x] 4.3 Add fzf requirement check at startup with helpful error message

## 5. Testing

- [x] 5.1 Test repository selection with multiple repos
- [x] 5.2 Test worktree selection with linked worktrees
- [x] 5.3 Test cancellation (Escape/Ctrl+C) behavior
- [x] 5.4 Test filtering functionality
- [x] 5.5 Test non-TTY environment handling
- [x] 5.6 Test fzf not installed error message
- [x] 5.7 Run all existing tests to ensure no regressions

## 6. Documentation

- [x] 6.1 Update README.md to mention fzf requirement
- [x] 6.2 Add fzf installation instructions for common platforms
- [x] 6.3 Document that FZF_DEFAULT_OPTS is respected
