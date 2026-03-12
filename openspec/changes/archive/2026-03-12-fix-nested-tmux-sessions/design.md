## Context

The `tm` tool manages tmux sessions for developers. Currently, when a user runs `tm` from inside an existing tmux session, the tmux command fails with the error message: `sessions should be nested with care, unset $TMUX to force`. This is because tmux, by default, prevents nested sessions to avoid complexity and potential confusion.

The `$TMUX` environment variable is set when inside a tmux session. When this variable is present, tmux commands that would create or attach to sessions fail unless forced. The current implementation doesn't account for this scenario.

## Goals / Non-Goals

**Goals:**
- Detect when `tm` is executed from within an existing tmux session
- Use `tmux switch-client` to seamlessly transition to the target session
- Maintain backward compatibility - users outside tmux sessions should see no change

**Non-Goals:**
- Supporting actual nested tmux sessions (sessions within sessions)
- Modifying tmux's default behavior regarding nested sessions
- Complex session management workflows beyond the existing `tm` functionality

## Decisions

### Decision: Detection Method
**Choice**: Use `$TMUX` environment variable check
**Rationale**: The `$TMUX` variable is the standard, reliable way to detect if currently inside a tmux session. It's set automatically by tmux and contains the path to the tmux socket.
**Alternative Considered**: Checking if `tmux list-clients` returns results. Rejected because it's slower and requires spawning a process.

### Decision: Use switch-client Command
**Choice**: When inside a tmux session, use `tmux switch-client -t <session>` to switch to the target session
**Rationale**: 
- `switch-client` is the canonical tmux command for changing sessions from within a client
- It works seamlessly without needing to detach or re-execute
- The tmux client naturally switches to the new session
- This is how other session managers like tmux-sessionizer handle this case
**Alternative Considered**: Detach and re-execute the tm command. Rejected because the detach kills the tm process before it can start the new session.

### Decision: Session Creation for switch-client
**Choice**: When the target session doesn't exist and we're inside tmux, create it detached first with `new-session -d`, then switch to it
**Rationale**: 
- `switch-client` requires the target session to already exist
- Creating with `-d` (detached) allows us to create without attempting to attach
- After creation, `switch-client` seamlessly transitions to the new session

### Decision: Implementation Location
**Choice**: Handle nested session logic within `AttachOrCreate` method
**Rationale**: 
- Centralizes session management logic in one place
- Keeps the command flow simple - no need for early returns or re-execution
- Easy to test and reason about

## Risks / Trade-offs

**Risk**: Edge cases where `$TMUX` is set but user isn't actually in an interactive session
**Mitigation**: Check if stdin is a TTY and `$TMUX` is set before triggering the switch-client logic. This filters out script/non-interactive usage.

## Migration Plan

No migration needed. This is a backward-compatible enhancement that fixes an existing limitation. Users currently experiencing the nested session error will see it resolved automatically after the update.
