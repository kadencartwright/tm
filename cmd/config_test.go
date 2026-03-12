package cmd

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"tm/internal/config"
	"tm/internal/discovery"
	"tm/internal/tmux"
	"tm/internal/worktree"
)

func TestConfigSetSearchPathCreatesConfigOnFirstWrite(t *testing.T) {
	root := t.TempDir()
	searchPath := filepath.Join(root, "repos")
	if err := os.MkdirAll(searchPath, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}

	store := config.NewTestStore(root)
	var stdout bytes.Buffer
	cmd := NewRootCmd(Dependencies{
		ConfigStore:  store,
		Discoverer:   &discovery.Service{},
		Selector:     &fakeSelector{},
		Inspector:    worktree.NewInspector(func(string) ([]byte, error) { return nil, nil }),
		TmuxLauncher: tmux.NewLauncher(&fakeTmuxCommander{}, nil, &stdout, &stdout),
		IsTTY:        func() bool { return true },
		Stdout:       &stdout,
		Stderr:       &stdout,
	})
	cmd.SetArgs([]string{"config", "set", "search-path", searchPath})

	if err := cmd.Execute(); err != nil {
		t.Fatalf("execute: %v", err)
	}

	cfg, err := store.Load()
	if err != nil {
		t.Fatalf("load: %v", err)
	}
	if cfg.SearchPath != searchPath {
		t.Fatalf("expected search path %q, got %q", searchPath, cfg.SearchPath)
	}
}

func TestConfigSetSearchPathRejectsInvalidPathWithoutOverwrite(t *testing.T) {
	root := t.TempDir()
	validPath := filepath.Join(root, "repos")
	if err := os.MkdirAll(validPath, 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}

	store := config.NewTestStore(root)
	if err := store.Save(config.Config{SearchPath: validPath}); err != nil {
		t.Fatalf("save: %v", err)
	}

	var stdout bytes.Buffer
	cmd := NewRootCmd(Dependencies{
		ConfigStore:  store,
		Discoverer:   &discovery.Service{},
		Selector:     &fakeSelector{},
		Inspector:    worktree.NewInspector(func(string) ([]byte, error) { return nil, nil }),
		TmuxLauncher: tmux.NewLauncher(&fakeTmuxCommander{}, nil, &stdout, &stdout),
		IsTTY:        func() bool { return true },
		Stdout:       &stdout,
		Stderr:       &stdout,
	})
	cmd.SetArgs([]string{"config", "set", "search-path", filepath.Join(root, "missing")})

	err := cmd.Execute()
	if err == nil {
		t.Fatalf("expected error")
	}

	cfg, loadErr := store.Load()
	if loadErr != nil {
		t.Fatalf("load: %v", loadErr)
	}
	if cfg.SearchPath != validPath {
		t.Fatalf("expected original path %q, got %q", validPath, cfg.SearchPath)
	}
}
