## Why

When `tm` is executed from inside an existing tmux session, tmux fails with the error:
`sessions should be nested with care, unset $TMUX to force`. This prevents users from starting a new tm session while already inside tmux, which is a common workflow for developers who switch between projects. The tool should handle this gracefully by using `tmux switch-client` when running inside a tmux session.

## What Changes

- Add detection for `$TMUX` environment variable to identify when running inside an existing tmux session
- When inside a tmux session, use `tmux switch-client -t <session>` to switch to the target session
- For new sessions, create them detached first with `new-session -d`, then switch to them
- Ensure the user is seamlessly transitioned to the new session without manual intervention

## Capabilities

### New Capabilities
- `nested-session-handling`: Detect when tm is run from inside a tmux session and use `switch-client` to seamlessly transition to the target session

### Modified Capabilities
<!-- No existing spec requirements are changing, only implementation details -->

## Impact

- **Command execution flow**: The `tm` command will check for `$TMUX` presence and use `switch-client` when appropriate
- **User experience**: Users can now run `tm` from within any tmux session without errors
- **No breaking changes**: This is a purely additive enhancement that fixes an existing limitation
