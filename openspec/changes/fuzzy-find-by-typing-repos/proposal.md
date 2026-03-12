## Why

The current implementation requires users to press '/' before typing to filter repos, which follows vim conventions but is less intuitive for users expecting modern fuzzy-finding behavior (like fzf or VS Code quick open). We will switch to direct typing for filtering, matching modern fuzzy finder UX.

## What Changes

- **Modify `selector` component** to enable immediate typing-to-filter mode without requiring '/' prefix
- **Remove vim-style keybindings** - '/' will no longer trigger filter mode
- **Navigation changes**: Arrow keys remain for navigation; 'j'/'k' will now type into the filter instead of navigating
- **UI improvements**: Reduce vertical padding between list items to show more repos at once

## Capabilities

### New Capabilities

- `fuzzy-find-typing`: Direct typing to filter repos without requiring '/' keypress

### Modified Capabilities

- `repo-session-selection`: Navigation behavior changes - 'j'/'k' keys now filter instead of navigate

## Impact

- **Affected**: `internal/selector/selector.go` - bubbletea list configuration and key handling
- **Dependencies**: charmbracelet/bubbles/list (already in use)
- **Breaking**: Users who relied on 'j'/'k' for navigation must use arrow keys instead
