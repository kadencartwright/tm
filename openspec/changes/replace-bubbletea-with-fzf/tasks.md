## 1. Dependency Cleanup

- [ ] 1.1 Remove bubbletea and bubbles from go.mod
- [ ] 1.2 Run `go mod tidy` to clean up unused dependencies
- [ ] 1.3 Verify no other files import bubbletea packages

## 2. FzfSelector Implementation

- [ ] 2.1 Create `FzfSelector` struct with stdin/stdout configuration
- [ ] 2.2 Implement fzf binary detection (`which fzf` equivalent)
- [ ] 2.3 Implement choice formatting (tab-separated values: Label\tDetails\tValue)
- [ ] 2.4 Implement fzf execution with exec.Command
- [ ] 2.5 Handle fzf exit codes (0=success, 1=no match, 130=cancelled)
- [ ] 2.6 Parse fzf output to return selected Choice
- [ ] 2.7 Set fzf options (--height=50%, --reverse, --border)
- [ ] 2.8 Pass choices via stdin to fzf
- [ ] 2.9 Return clear error if fzf not installed

## 3. Interface Compatibility

- [ ] 3.1 Keep `Choice` struct unchanged (Label, Details, Value, Title(), Description(), FilterValue())
- [ ] 3.2 Keep `Select(title string, items []Choice) (Choice, bool, error)` signature
- [ ] 3.3 Keep `RepoChoices()` and `TargetChoices()` helper functions
- [ ] 3.4 Keep `IsTTY()` function for TTY detection
- [ ] 3.5 Rename `NewBubbleSelector` to `NewFzfSelector` or keep for compatibility

## 4. Integration Updates

- [ ] 4.1 Update `cmd/root.go` to use `NewFzfSelector` instead of `NewBubbleSelector`
- [ ] 4.2 Verify `cmd/root_test.go` fake selector still works (should use interface)
- [ ] 4.3 Add fzf requirement check at startup with helpful error message

## 5. Testing

- [ ] 5.1 Test repository selection with multiple repos
- [ ] 5.2 Test worktree selection with linked worktrees
- [ ] 5.3 Test cancellation (Escape/Ctrl+C) behavior
- [ ] 5.4 Test filtering functionality
- [ ] 5.5 Test non-TTY environment handling
- [ ] 5.6 Test fzf not installed error message
- [ ] 5.7 Run all existing tests to ensure no regressions

## 6. Documentation

- [ ] 6.1 Update README.md to mention fzf requirement
- [ ] 6.2 Add fzf installation instructions for common platforms
- [ ] 6.3 Document that FZF_DEFAULT_OPTS is respected
