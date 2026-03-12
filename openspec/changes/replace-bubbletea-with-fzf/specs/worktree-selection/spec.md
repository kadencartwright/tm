## MODIFIED Requirements

### Requirement: Repositories with multiple checkout targets show a second selector
After a repository is selected, the CLI MUST inspect available checkout targets and show a second fzf selector when the repository has more than one available target.

**Change**: Changed from "Bubble Tea selector" to "fzf selector".

#### Scenario: Repository has linked worktrees
- **WHEN** a user selects a repository whose Git metadata reports linked worktrees
- **THEN** the CLI opens a second fzf selector containing the main checkout and each linked worktree

#### Scenario: Repository has only main checkout
- **WHEN** a user selects a repository that has no linked worktrees
- **THEN** the CLI skips the second selector and continues with the main checkout path

### Requirement: Worktree list shows name and full path
The worktree selector MUST display each available checkout target with a concise name and full path description.

**Change**: No change in requirements - implementation now uses fzf's display format.

#### Scenario: Render main checkout target
- **WHEN** the main checkout is included in the second selector
- **THEN** it appears as a selectable item with a readable title and its full path as the description

#### Scenario: Render linked worktree target
- **WHEN** a linked worktree is included in the second selector
- **THEN** it appears as a selectable item with a readable title and its full path as the description
