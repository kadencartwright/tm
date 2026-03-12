package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestConfigPathUsesXDGConfigHome(t *testing.T) {
	root := t.TempDir()
	store := NewTestStore(root)
	path, err := store.ConfigPath()
	if err != nil {
		t.Fatalf("config path: %v", err)
	}
	want := filepath.Join(root, ".config", "tm", "config.toml")
	if path != want {
		t.Fatalf("expected %q, got %q", want, path)
	}
}

func TestConfigPathFallsBackToHomeConfigDir(t *testing.T) {
	root := t.TempDir()
	store := &Store{
		getenv:      func(string) string { return "" },
		userHomeDir: func() (string, error) { return root, nil },
	}

	path, err := store.ConfigPath()
	if err != nil {
		t.Fatalf("config path: %v", err)
	}
	want := filepath.Join(root, ".config", "tm", "config.toml")
	if path != want {
		t.Fatalf("expected %q, got %q", want, path)
	}
}

func TestLoadCreatesMissingConfig(t *testing.T) {
	root := t.TempDir()
	store := NewTestStore(root)

	cfg, err := store.Load()
	if err != nil {
		t.Fatalf("load: %v", err)
	}
	if cfg.SearchPath != "" {
		t.Fatalf("expected empty config, got %+v", cfg)
	}

	path, err := store.ConfigPath()
	if err != nil {
		t.Fatalf("config path: %v", err)
	}
	if _, err := os.Stat(path); err != nil {
		t.Fatalf("expected created config file: %v", err)
	}
}

func TestSaveUpdatesConfig(t *testing.T) {
	root := t.TempDir()
	searchPath := filepath.Join(root, "repos")
	if err := os.MkdirAll(searchPath, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}

	store := NewTestStore(root)
	if err := store.Save(Config{SearchPath: searchPath}); err != nil {
		t.Fatalf("save: %v", err)
	}

	cfg, err := store.Load()
	if err != nil {
		t.Fatalf("load: %v", err)
	}
	if cfg.SearchPath != searchPath {
		t.Fatalf("expected %q, got %q", searchPath, cfg.SearchPath)
	}
}
