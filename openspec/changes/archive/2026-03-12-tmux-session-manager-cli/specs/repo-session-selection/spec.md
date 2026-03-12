## ADDED Requirements

### Requirement: Default command shows repository selector
The CLI MUST invoke an interactive Bubble Tea repository selector when `tm` is run without subcommands.

#### Scenario: Launch selector with configured path
- **WHEN** a user runs `tm` and a valid `search_path` is configured
- **THEN** the CLI displays a Bubble Tea list populated with repositories discovered from that path

#### Scenario: Handle non-interactive execution
- **WHEN** a user runs `tm` in a non-TTY environment
- **THEN** the CLI MUST exit with a clear error indicating that interactive selection requires a TTY

### Requirement: Repository list follows simple title-and-description layout
The repository selector MUST follow the `list-simple` Bubble Tea interaction style and display each repository with a name title and full path description.

#### Scenario: Render repository list item
- **WHEN** a repository is displayed in the selector
- **THEN** the item shows the repository name as the title and the full filesystem path as the description

#### Scenario: Filter repository list
- **WHEN** a user types into the selector filter
- **THEN** the list narrows visible repositories using Bubble Tea list filtering behavior

### Requirement: Repository discovery identifies code repositories
The CLI MUST discover candidate repositories by scanning directories under the configured search path and selecting directories that contain Git metadata.

#### Scenario: Discover git repositories in child directories
- **WHEN** the configured search path contains child directories where some include a `.git` directory
- **THEN** only directories containing `.git` are included in the selector list

#### Scenario: Ignore non-repository directories
- **WHEN** child directories exist without Git metadata
- **THEN** those directories are excluded from selector results

### Requirement: Selection returns repository context for tmux workflow
After a repository is selected, the CLI MUST produce repository context that downstream tmux workflow commands can consume.

#### Scenario: Successful repository selection
- **WHEN** a user chooses a repository from the fuzzy selector
- **THEN** the CLI returns the selected repository path and name as the active context for subsequent tmux actions

#### Scenario: User cancels selection
- **WHEN** a user exits the selector without choosing a repository
- **THEN** the CLI exits gracefully without launching tmux actions
