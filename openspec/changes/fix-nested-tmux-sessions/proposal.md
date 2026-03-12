## Why

When `tm` is executed from inside an existing tmux session, tmux fails with the error:
`sessions should be nested with care, unset $TMUX to force`. This prevents users from starting a new tm session while already inside tmux, which is a common workflow for developers who switch between projects. The tool should handle this gracefully by detecting the nested session situation and detaching before attempting to create or attach to a new session.

## What Changes

- Add detection for `$TMUX` environment variable to identify when running inside an existing tmux session
- Implement automatic detachment from the current tmux session before executing tmux commands
- Execute tmux session creation/attachment commands from outside the nested session context
- Ensure the user is seamlessly transitioned to the new session without manual intervention

## Capabilities

### New Capabilities
- `nested-session-handling`: Detect when tm is run from inside a tmux session and handle the situation gracefully by detaching and re-executing outside the session context

### Modified Capabilities
<!-- No existing spec requirements are changing, only implementation details -->

## Impact

- **Command execution flow**: The `tm` command wrapper will check for `$TMUX` presence and potentially detach/re-execute
- **User experience**: Users can now run `tm` from within any tmux session without errors
- **No breaking changes**: This is a purely additive enhancement that fixes an existing limitation
