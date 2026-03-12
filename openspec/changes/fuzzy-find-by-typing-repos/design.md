## Context

The current `tm` tool uses the charmbracelet/bubbletea framework with the bubbles/list component for repo selection. The list component defaults to vim-style keybindings where users must press '/' to enter filter mode before typing search terms. We are changing to direct-typing behavior, matching modern fuzzy finders like fzf.

## Goals / Non-Goals

**Goals:**
- Enable direct typing-to-filter behavior in the repo selector
- Remove the '/' key requirement to enter filter mode
- Keep navigation simple with arrow keys only

**Non-Goals:**
- Maintain vim-style keybindings as an option
- Support custom key remapping
- Add filtering to other components (out of scope for this change)

## Decisions

**1. Implementation Approach**
- Use bubbletea's `KeyMap` customization to enable direct filtering
- The list component supports setting `KeyMap` to customize filter behavior
- Enable `SetFilteringEnabled(true)` and customize the filter input handling

**2. Navigation Strategy**
- Arrow keys (Up/Down) for navigation
- 'j' and 'k' will type into the filter instead of navigating
- This is **BREAKING** for users accustomed to vim navigation

**3. Simplicity**
- No configuration option - this is a one-time behavioral change
- Cleaner codebase without mode switching logic
