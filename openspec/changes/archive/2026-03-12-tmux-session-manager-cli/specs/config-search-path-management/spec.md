## ADDED Requirements

### Requirement: Config command sets search path
The CLI MUST provide `tm config set search-path <path>` to store the repository discovery root path.

#### Scenario: Set search path successfully
- **WHEN** a user runs `tm config set search-path ~/code`
- **THEN** the command validates and persists the path in the config file

#### Scenario: Reject invalid path input
- **WHEN** a user runs `tm config set search-path <path>` with an invalid or non-existent path
- **THEN** the command fails with a clear validation error and does not overwrite existing valid config

### Requirement: Config file is created on first write
The CLI MUST create the config directory and TOML config file if they do not exist when executing `config set`.

#### Scenario: Create missing config directory and file
- **WHEN** `~/.config/tm/` and `config.toml` are absent and user runs `tm config set search-path <path>`
- **THEN** the CLI creates the directory and file and writes the `search_path` value

#### Scenario: Update existing config file
- **WHEN** a config file already exists and user runs `tm config set search-path <path>`
- **THEN** the CLI updates only the `search_path` field while preserving valid TOML structure

### Requirement: Config is loaded for default command behavior
The CLI MUST load configured `search_path` before repository discovery begins.

#### Scenario: Missing search path configuration
- **WHEN** a user runs `tm` without configured `search_path`
- **THEN** the CLI returns a clear instruction to run `tm config set search-path <path>`

#### Scenario: Use stored search path for discovery
- **WHEN** a valid `search_path` is present in `config.toml`
- **THEN** `tm` uses that path as the source for repository discovery
