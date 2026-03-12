## ADDED Requirements

### Requirement: Detect nested tmux session
The system SHALL detect when it is running inside an existing tmux session by checking for the presence of the `$TMUX` environment variable.

#### Scenario: TMUX variable is set
- **WHEN** the `$TMUX` environment variable is set and non-empty
- **THEN** the system SHALL identify this as a nested session scenario

#### Scenario: TMUX variable is not set
- **WHEN** the `$TMUX` environment variable is unset or empty
- **THEN** the system SHALL proceed with normal execution

### Requirement: Detach from current session
When running inside a tmux session, the system SHALL detach from the current session before executing any tmux commands that would create or attach to sessions.

#### Scenario: Detach from existing session
- **GIVEN** the user is currently attached to a tmux session
- **WHEN** the user executes `tm` with any valid arguments
- **THEN** the system SHALL first detach from the current session using `tmux detach-client`
- **AND** the system SHALL then proceed to execute the requested tm command from outside the tmux session context

#### Scenario: Detach and switch to new session
- **GIVEN** the user is currently attached to a tmux session named "old-project"
- **WHEN** the user executes `tm new-project`
- **THEN** the system SHALL detach from "old-project"
- **AND** the system SHALL attach the user to the "new-project" session
- **AND** the user SHALL see no error messages about nested sessions

### Requirement: Preserve command arguments
When detaching and re-executing, the system SHALL preserve all original command-line arguments passed to `tm`.

#### Scenario: Arguments preserved during re-execution
- **GIVEN** the user executes `tm --list --format=json` from inside a tmux session
- **WHEN** the system handles the nested session scenario
- **THEN** the same arguments (`--list --format=json`) SHALL be passed to the re-executed command
- **AND** the output SHALL be identical to running the command outside a tmux session
