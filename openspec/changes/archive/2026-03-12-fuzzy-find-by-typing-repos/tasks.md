## 1. Selector Component Updates

- [x] 1.1 Enable direct typing in bubbletea list: modify `internal/selector/selector.go` to handle character keys as filter input
- [x] 1.2 Update `Update` method to route printable characters to the filter input instead of treating them as commands
- [x] 1.3 Ensure arrow keys (Up/Down) still navigate the list
- [x] 1.4 Remove or disable any vim-specific navigation key handling for 'j'/'k'
- [x] 1.5 Preserve existing behavior for Enter, Escape, 'q', Ctrl+C
- [x] 1.6 Reduce padding between list items in the bubbletea list delegate (tighten the UI)
- [x] 1.7 Upgrade `charmbracelet/bubbles` from v0.20.0 to v0.21.0 to access `SetFilterState` method for starting in filter mode

## 2. Testing

- [x] 2.1 Update selector tests in `internal/selector/selector_test.go` to reflect new behavior
- [x] 2.2 Test scenario: Direct typing filters repos
- [x] 2.3 Test scenario: Arrow keys navigate the list
- [x] 2.4 Test scenario: 'j'/'k' characters now filter instead of navigate
- [x] 2.5 Test scenario: Enter selects repo
- [x] 2.6 Test scenario: Escape/q/Ctrl+C cancels selection

## 3. Documentation

- [x] 3.1 Update README.md to document the new behavior
- [x] 3.2 Document that users should use arrow keys instead of j/k for navigation
- [x] 3.3 Add a note about the **BREAKING** change for existing vim users
