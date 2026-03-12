## Context

The `tm` tool manages tmux sessions for developers. Currently, when a user runs `tm` from inside an existing tmux session, the tmux command fails with the error message: `sessions should be nested with care, unset $TMUX to force`. This is because tmux, by default, prevents nested sessions to avoid complexity and potential confusion.

The `$TMUX` environment variable is set when inside a tmux session. When this variable is present, tmux commands that would create or attach to sessions fail unless forced. The current implementation doesn't account for this scenario.

## Goals / Non-Goals

**Goals:**
- Detect when `tm` is executed from within an existing tmux session
- Automatically detach from the current tmux session before executing tmux commands
- Seamlessly transition the user to the new/desired tmux session without manual intervention
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

### Decision: Detach and Re-execute Strategy
**Choice**: Detach current session, then re-execute the tm command from a non-tmux context
**Rationale**: 
- Clean approach that respects tmux's design intention (no nested sessions)
- User ends up in the desired session seamlessly
- Avoids complex shell manipulation
**Alternative Considered**: Using `tmux -L` with a different socket name. Rejected because it creates parallel session namespaces that could confuse users.

### Decision: Implementation Location
**Choice**: Early in the `tm` command execution, before any tmux commands are run
**Rationale**: 
- Prevents any tmux commands from failing due to nested session restrictions
- Centralizes the handling logic in one place
- Fail-fast approach - detect and handle immediately

## Risks / Trade-offs

**Risk**: Detaching unexpectedly might disrupt the user's workflow if they didn't intend to leave the current session
**Mitigation**: The user explicitly ran `tm` to switch sessions, so detaching aligns with their intent. Document this behavior in help text.

**Risk**: Re-execution might lose environment variables or context from the original shell
**Mitigation**: The new session inherits the parent shell environment. Any tmux-specific environment (like `$TMUX_PANE`) is naturally cleared by the detach, which is the desired behavior.

**Risk**: Edge cases where `$TMUX` is set but user isn't actually in an interactive session
**Mitigation**: Check if stdin is a TTY and `$TMUX` is set before triggering the detach logic. This filters out script/non-interactive usage.

## Migration Plan

No migration needed. This is a backward-compatible enhancement that fixes an existing limitation. Users currently experiencing the nested session error will see it resolved automatically after the update.
