## ADDED Requirements

### Requirement: Detect nested tmux session
The system SHALL detect when it is running inside an existing tmux session by checking for the presence of the `$TMUX` environment variable.

#### Scenario: TMUX variable is set
- **WHEN** the `$TMUX` environment variable is set and non-empty
- **THEN** the system SHALL identify this as a nested session scenario

#### Scenario: TMUX variable is not set
- **WHEN** the `$TMUX` environment variable is unset or empty
- **THEN** the system SHALL proceed with normal execution

### Requirement: Use switch-client for nested sessions
When running inside a tmux session, the system SHALL use `tmux switch-client` to change to the target session instead of attempting to create/attach directly.

#### Scenario: Switch to existing session from within tmux
- **GIVEN** the user is currently attached to a tmux session named "current-session"
- **AND** a session named "target-session" already exists
- **WHEN** the user executes `tm` and selects "target-session"
- **THEN** the system SHALL execute `tmux switch-client -t target-session`
- **AND** the user SHALL be switched to "target-session" without errors

#### Scenario: Create and switch to new session from within tmux
- **GIVEN** the user is currently attached to a tmux session named "current-session"
- **AND** a session named "new-project" does not exist
- **WHEN** the user executes `tm` and selects "new-project"
- **THEN** the system SHALL first execute `tmux new-session -d -s new-project -c <path>` to create the session detached
- **AND** then execute `tmux switch-client -t new-project` to switch to it
- **AND** the user SHALL be switched to "new-project" without errors

### Requirement: TTY check for nested session handling
The system SHALL only use `switch-client` behavior when stdin is a TTY to avoid breaking non-interactive usage.

#### Scenario: Nested session with TTY
- **GIVEN** the `$TMUX` environment variable is set
- **AND** stdin is a TTY
- **WHEN** the user executes `tm`
- **THEN** the system SHALL use `switch-client` behavior

#### Scenario: Nested session without TTY
- **GIVEN** the `$TMUX` environment variable is set
- **AND** stdin is not a TTY
- **WHEN** a script executes `tm`
- **THEN** the system SHALL fall back to standard behavior
