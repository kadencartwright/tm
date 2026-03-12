package config

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/BurntSushi/toml"
)

type Config struct {
	SearchPath string `toml:"search_path"`
}

type Store struct {
	getenv      func(string) string
	userHomeDir func() (string, error)
}

func NewStore() *Store {
	return &Store{
		getenv:      os.Getenv,
		userHomeDir: os.UserHomeDir,
	}
}

func NewTestStore(root string) *Store {
	return &Store{
		getenv: func(key string) string {
			if key == "XDG_CONFIG_HOME" {
				return filepath.Join(root, ".config")
			}
			return ""
		},
		userHomeDir: func() (string, error) { return root, nil },
	}
}

func (s *Store) ConfigPath() (string, error) {
	base := strings.TrimSpace(s.getenv("XDG_CONFIG_HOME"))
	if base == "" {
		home, err := s.userHomeDir()
		if err != nil {
			return "", fmt.Errorf("resolve home directory: %w", err)
		}
		base = filepath.Join(home, ".config")
	}

	return filepath.Join(base, "tm", "config.toml"), nil
}

func (s *Store) Load() (Config, error) {
	path, err := s.ConfigPath()
	if err != nil {
		return Config{}, err
	}
	if err := ensureConfigFile(path); err != nil {
		return Config{}, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("read config: %w", err)
	}
	if len(bytes.TrimSpace(data)) == 0 {
		return Config{}, nil
	}

	var cfg Config
	if _, err := toml.Decode(string(data), &cfg); err != nil {
		return Config{}, fmt.Errorf("decode config: %w", err)
	}

	return cfg, nil
}

func (s *Store) Save(cfg Config) error {
	path, err := s.ConfigPath()
	if err != nil {
		return err
	}
	if err := ensureConfigFile(path); err != nil {
		return err
	}

	var buf bytes.Buffer
	if err := toml.NewEncoder(&buf).Encode(cfg); err != nil {
		return fmt.Errorf("encode config: %w", err)
	}

	if err := os.WriteFile(path, buf.Bytes(), 0o644); err != nil {
		return fmt.Errorf("write config: %w", err)
	}

	return nil
}

func (s *Store) NormalizeSearchPath(path string) (string, error) {
	normalized, err := expandPath(path, s.userHomeDir)
	if err != nil {
		return "", err
	}
	info, err := os.Stat(normalized)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("search path %q does not exist", normalized)
		}
		return "", fmt.Errorf("validate search path: %w", err)
	}
	if !info.IsDir() {
		return "", fmt.Errorf("search path %q is not a directory", normalized)
	}

	return filepath.Clean(normalized), nil
}

func ensureConfigFile(path string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return fmt.Errorf("create config directory: %w", err)
	}

	if _, err := os.Stat(path); err == nil {
		return nil
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("stat config: %w", err)
	}

	if err := os.WriteFile(path, nil, 0o644); err != nil {
		return fmt.Errorf("create config file: %w", err)
	}

	return nil
}

func expandPath(path string, homeDir func() (string, error)) (string, error) {
	trimmed := strings.TrimSpace(path)
	if trimmed == "" {
		return "", fmt.Errorf("search path cannot be empty")
	}

	if trimmed == "~" || strings.HasPrefix(trimmed, "~/") {
		home, err := homeDir()
		if err != nil {
			return "", fmt.Errorf("resolve home directory: %w", err)
		}
		if trimmed == "~" {
			return home, nil
		}
		trimmed = filepath.Join(home, trimmed[2:])
	}

	if !filepath.IsAbs(trimmed) {
		abs, err := filepath.Abs(trimmed)
		if err != nil {
			return "", fmt.Errorf("resolve search path: %w", err)
		}
		trimmed = abs
	}

	return filepath.Clean(trimmed), nil
}
