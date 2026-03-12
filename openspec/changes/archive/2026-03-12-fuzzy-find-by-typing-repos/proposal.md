## Why

The current implementation requires users to press '/' before typing to filter repos, which follows vim conventions but is less intuitive for users expecting modern fuzzy-finding behavior (like fzf or VS Code quick open). We will switch to direct typing for filtering with the filter input always visible, matching modern fuzzy finder UX.

## What Changes

- **Modify `selector` component** to show filter input immediately on open (always-visible like fzf)
- **Enable immediate typing-to-filter** - all character keys go directly to the filter
- **Remove vim-style keybindings completely** - '/' will no longer trigger filter mode (not needed since filter is always visible)
- **Remove 'j'/'k' navigation** - character keys now filter; use Up/Down arrows for navigation
- **UI improvements**: Reduce vertical padding between list items to show more repos at once

## Capabilities

### New Capabilities

- `fuzzy-find-typing`: Filter input always visible; direct typing filters repos immediately

### Modified Capabilities

- `repo-session-selection`: Navigation behavior changes - 'j'/'k' keys removed, filter always visible

## Impact

- **Affected**: `internal/selector/selector.go` - bubbletea list configuration and key handling
- **Dependencies**: charmbracelet/bubbles/list (already in use)
- **Breaking**: Users who relied on 'j'/'k' for navigation must use arrow keys instead; no backwards compatibility for vim-style navigation
