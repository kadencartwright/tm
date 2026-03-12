## ADDED Requirements

### Requirement: Repositories with multiple checkout targets show a second selector
After a repository is selected, the CLI MUST inspect available checkout targets and show a second Bubble Tea selector when the repository has more than one available target.

#### Scenario: Repository has linked worktrees
- **WHEN** a user selects a repository whose Git metadata reports linked worktrees
- **THEN** the CLI opens a second selector containing the main checkout and each linked worktree

#### Scenario: Repository has only main checkout
- **WHEN** a user selects a repository that has no linked worktrees
- **THEN** the CLI skips the second selector and continues with the main checkout path

### Requirement: Worktree list shows name and full path
The worktree selector MUST display each available checkout target with a concise name and full path description.

#### Scenario: Render main checkout target
- **WHEN** the main checkout is included in the second selector
- **THEN** it appears as a selectable item with a readable title and its full path as the description

#### Scenario: Render linked worktree target
- **WHEN** a linked worktree is included in the second selector
- **THEN** it appears as a selectable item with a readable title and its full path as the description

### Requirement: Worktree discovery uses Git metadata
The CLI MUST derive worktree choices from Git worktree metadata rather than filesystem guessing.

#### Scenario: Read Git worktree list
- **WHEN** the CLI prepares worktree choices for a selected repository
- **THEN** it uses Git worktree metadata to enumerate the main checkout and linked worktrees

#### Scenario: Handle worktree inspection failure
- **WHEN** Git worktree metadata cannot be read for the selected repository
- **THEN** the CLI returns a clear error and does not launch tmux for an unknown target
