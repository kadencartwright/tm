## 1. Setup and Detection

- [x] 1.1 Add `IsNestedSession() bool` function that checks if `$TMUX` env var is set
- [x] 1.2 Add `IsTerminal() bool` function to check if stdin is a TTY
- [x] 1.3 Add `SessionExists(session string) bool` method to check if a tmux session exists

## 2. Core Implementation

- [x] 2.1 Modify `AttachOrCreate` in `internal/tmux/tmux.go` to check `IsNestedSession()` and `IsTerminal()`
- [x] 2.2 When nested, create session detached first using `new-session -d` if it doesn't exist
- [x] 2.3 When nested, use `switch-client -t <session>` to switch to the target session
- [x] 2.4 When not nested, use standard `new-session -A` behavior

## 3. Testing

- [x] 3.1 Add unit tests for `IsNestedSession()` with `$TMUX` set and unset scenarios
- [x] 3.2 Add unit tests for `SessionExists()` method
- [x] 3.3 Update test mocks to implement the Commander interface
- [x] 3.4 Add tests for the nested session code path

## 4. Documentation

- [x] 4.1 Update help text to mention the automatic session switching behavior when running inside tmux
- [x] 4.2 Add note to README about nested session handling
- [x] 4.3 Update CHANGELOG with the fix description
