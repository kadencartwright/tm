## ADDED Requirements

### Requirement: Selected target opens in tmux
After the user selects a repository or worktree target, the CLI MUST attach to an existing tmux session for that target or create one if none exists.

#### Scenario: Existing session found
- **WHEN** the selected target already has a matching tmux session
- **THEN** the CLI attaches the user to that existing session

#### Scenario: No session exists
- **WHEN** the selected target does not have a matching tmux session
- **THEN** the CLI creates a new tmux session rooted at the selected target path and attaches to it

### Requirement: Session identity is deterministic per selected target
The CLI MUST derive tmux session identity from the selected repository or worktree target so repeated launches resolve to the same session.

#### Scenario: Reopen same repository target
- **WHEN** a user selects the same repository target on two different launches
- **THEN** both launches resolve to the same tmux session identity

#### Scenario: Distinguish different worktrees
- **WHEN** a user selects different worktree targets for the same repository
- **THEN** each worktree resolves to a distinct tmux session identity

### Requirement: Tmux errors are reported clearly
The CLI MUST surface tmux failures with actionable error messaging.

#### Scenario: Tmux is unavailable
- **WHEN** the `tmux` executable is not installed or cannot be executed
- **THEN** the CLI exits with an error explaining that tmux is required

#### Scenario: Session create or attach fails
- **WHEN** tmux returns an error while creating or attaching to a session
- **THEN** the CLI exits with the tmux failure details and does not report success
