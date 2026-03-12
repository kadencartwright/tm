## 1. Setup and Detection

- [x] 1.1 Add `IsNestedSession() bool` method to tmux Launcher that checks if `$TMUX` env var is set
- [x] 1.2 Add `DetachClient() error` method to tmux Commander interface and ExecCommander implementation
- [x] 1.3 Add `ReExecute(args []string) error` helper function to handle command re-execution outside tmux context

## 2. Core Implementation

- [x] 2.1 Modify `AttachOrCreate` in `internal/tmux/tmux.go` to check `IsNestedSession()` at the start
- [x] 2.2 Implement detach logic: if nested, call `DetachClient()` before executing tmux new-session
- [x] 2.3 Ensure command arguments are preserved when re-executing after detach
- [x] 2.4 Add TTY check alongside `$TMUX` detection to avoid triggering in non-interactive scripts

## 3. Testing

- [x] 3.1 Add unit tests for `IsNestedSession()` with `$TMUX` set and unset scenarios
- [x] 3.2 Add mock implementation of `DetachClient()` in test doubles
- [x] 3.3 Add integration test verifying the full detach-and-switch workflow
- [x] 3.4 Test that arguments are preserved through the detach/re-execute flow

## 4. Documentation

- [x] 4.1 Update help text to mention the automatic detach behavior when running inside tmux
- [x] 4.2 Add note to README about nested session handling
- [x] 4.3 Update CHANGELOG with the fix description
