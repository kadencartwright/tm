## ADDED Requirements

### Requirement: Filter input visible by default
The system SHALL display the filter input field immediately when the repo selector opens, without requiring any keypress to activate it.

#### Scenario: Filter visible on open
- **WHEN** the repo selector is displayed
- **THEN** the filter input field SHALL be visible at the top of the list
- **AND** the user SHALL be able to type immediately to filter repos

### Requirement: Direct typing always filters
The system SHALL treat all printable character key presses as filter input at all times.

#### Scenario: User types to filter repos
- **WHEN** the repo selector is displayed
- **AND** the user types "myrepo"
- **THEN** those characters SHALL appear in the filter input
- **AND** the repo list SHALL be filtered to show only repos matching "myrepo"

#### Scenario: Filter maintains navigation capability
- **WHEN** the repo selector is displayed
- **AND** the user presses Up/Down arrow keys
- **THEN** the selection SHALL move up/down in the filtered list

### Requirement: Compact list item display
The system SHALL display list items with minimal vertical padding to maximize the number of visible items.

#### Scenario: Tight list spacing
- **WHEN** the repo selector is displayed
- **THEN** each list item SHALL have reduced vertical padding (spacing)
- **AND** more repos SHALL be visible in the viewport compared to the previous spacing

## REMOVED Requirements

### Requirement: Vim-style filter trigger
**Reason**: Filter input is always visible, no trigger key needed
**Migration**: Users can now type immediately to filter; arrow keys replace j/k for navigation

### Requirement: 'j'/'k' navigation
**Reason**: Character keys now filter instead of navigate
**Migration**: Use Up/Down arrow keys for navigation

## UNCHANGED Requirements

### Requirement: Navigation with arrow keys
The system SHALL preserve Up/Down arrow key navigation.

#### Scenario: Arrow key navigation
- **WHEN** the repo selector is displayed
- **AND** the user presses Down arrow 3 times
- **THEN** the 4th repo in the list SHALL be selected

#### Scenario: Arrow navigation in filtered list
- **WHEN** the repo selector is filtered
- **AND** the user presses Up/Down arrow keys
- **THEN** the selection SHALL move through the filtered results

### Requirement: Enter selects repo
The system SHALL allow pressing Enter to select the highlighted repo.

#### Scenario: Enter selects repo
- **WHEN** the user has navigated to a repo in the selector
- **AND** the user presses Enter
- **THEN** the selected repo SHALL be returned as the choice
- **AND** the selector SHALL close

### Requirement: Escape/q/Ctrl+C cancels selection
The system SHALL allow pressing Escape, 'q', or Ctrl+C to cancel selection.

#### Scenario: Cancel selection
- **WHEN** the repo selector is displayed
- **AND** the user presses Escape, 'q', or Ctrl+C
- **THEN** the selector SHALL close without returning a selection
