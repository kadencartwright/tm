## MODIFIED Requirements

### Requirement: Default command shows repository selector
The CLI MUST invoke an interactive fzf repository selector when `tm` is run without subcommands.

**Change**: Changed from "Bubble Tea" to "fzf" as the selection mechanism.

#### Scenario: Launch selector with configured path
- **WHEN** a user runs `tm` and a valid `search_path` is configured
- **THEN** the CLI displays an fzf list populated with repositories discovered from that path

#### Scenario: Handle non-interactive execution
- **WHEN** a user runs `tm` in a non-TTY environment
- **THEN** the CLI MUST exit with a clear error indicating that interactive selection requires a TTY

### Requirement: Repository list follows simple title-and-description layout
The repository selector MUST follow the fzf layout and display each repository with a name title and full path description.

**Change**: Changed from "list-simple Bubble Tea interaction style" to "fzf layout".

#### Scenario: Render repository list item
- **WHEN** a repository is displayed in the selector
- **THEN** the item shows the repository name as the title and the full filesystem path as the description

#### Scenario: Filter repository list
- **WHEN** a user types into the selector filter
- **THEN** the list narrows visible repositories using fzf's fuzzy matching behavior

## ADDED Requirements

### Requirement: fzf binary availability check
The CLI SHALL verify that the fzf binary is available in PATH before attempting to launch the selector.

#### Scenario: fzf not installed
- **WHEN** a user runs `tm` without fzf installed
- **THEN** the CLI displays an error message indicating fzf is required
- **AND** the CLI provides installation instructions
