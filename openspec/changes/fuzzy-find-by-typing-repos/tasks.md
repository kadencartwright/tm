## 1. Selector Component Updates

- [ ] 1.1 Enable direct typing in bubbletea list: modify `internal/selector/selector.go` to handle character keys as filter input
- [ ] 1.2 Update `Update` method to route printable characters to the filter input instead of treating them as commands
- [ ] 1.3 Ensure arrow keys (Up/Down) still navigate the list
- [ ] 1.4 Remove or disable any vim-specific navigation key handling for 'j'/'k'
- [ ] 1.5 Preserve existing behavior for Enter, Escape, 'q', Ctrl+C
- [ ] 1.6 Reduce padding between list items in the bubbletea list delegate (tighten the UI)

## 2. Testing

- [ ] 2.1 Update selector tests in `internal/selector/selector_test.go` to reflect new behavior
- [ ] 2.2 Test scenario: Direct typing filters repos
- [ ] 2.3 Test scenario: Arrow keys navigate the list
- [ ] 2.4 Test scenario: 'j'/'k' characters now filter instead of navigate
- [ ] 2.5 Test scenario: Enter selects repo
- [ ] 2.6 Test scenario: Escape/q/Ctrl+C cancels selection

## 3. Documentation

- [ ] 3.1 Update README.md to document the new behavior
- [ ] 3.2 Document that users should use arrow keys instead of j/k for navigation
- [ ] 3.3 Add a note about the **BREAKING** change for existing vim users
