# Changelog

All notable changes to this project will be documented in this file.

## [Unreleased]

### Fixed

- Fixed nested tmux session error when running `tm` from inside an existing tmux session. The tool now automatically detaches from the current session before creating or attaching to a new one.

