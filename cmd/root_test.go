package cmd

import (
	"bytes"
	"io"
	"os"
	"path/filepath"
	"testing"

	"tm/internal/config"
	"tm/internal/discovery"
	"tm/internal/selector"
	"tm/internal/tmux"
	"tm/internal/worktree"
)

type fakeSelector struct {
	choices []selector.Choice
	ok      bool
	err     error
	titles  []string
	seen    [][]selector.Choice
}

func (f *fakeSelector) Select(title string, items []selector.Choice) (selector.Choice, bool, error) {
	f.titles = append(f.titles, title)
	f.seen = append(f.seen, items)
	if f.err != nil {
		return selector.Choice{}, false, f.err
	}
	return f.choices[len(f.titles)-1], f.ok, nil
}

type fakeTmuxCommander struct {
	last []string
	err  error
}

func (f *fakeTmuxCommander) LookPath(string) (string, error) { return "/usr/bin/tmux", nil }
func (f *fakeTmuxCommander) Run(args []string, dir string, _ io.Reader, _, _ io.Writer) error {
	f.last = append([]string{dir}, args...)
	return f.err
}
func (f *fakeTmuxCommander) DetachClient() error { return nil }

func TestRootCommandNonTTYError(t *testing.T) {
	root := t.TempDir()
	searchPath := filepath.Join(root, "repos")
	repoPath := filepath.Join(searchPath, "alpha")
	if err := os.MkdirAll(filepath.Join(repoPath, ".git"), 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}

	store := config.NewTestStore(root)
	if err := store.Save(config.Config{SearchPath: searchPath}); err != nil {
		t.Fatalf("save config: %v", err)
	}

	var stdout bytes.Buffer
	deps := Dependencies{
		ConfigStore:  store,
		Discoverer:   &discovery.Service{},
		Selector:     &fakeSelector{},
		Inspector:    worktree.NewInspector(func(string) ([]byte, error) { return nil, nil }),
		TmuxLauncher: tmux.NewLauncher(&fakeTmuxCommander{}, nil, &stdout, &stdout),
		IsTTY:        func() bool { return false },
		Stdout:       &stdout,
		Stderr:       &stdout,
	}

	err := NewRootCmd(deps).Execute()
	if err == nil || err.Error() != "interactive selection requires a TTY" {
		t.Fatalf("expected non-tty error, got %v", err)
	}
}

func TestRootCommandSelectionSuccessSkipsSecondSelectorForSingleTarget(t *testing.T) {
	root := t.TempDir()
	searchPath := filepath.Join(root, "repos")
	repoPath := filepath.Join(searchPath, "alpha")
	if err := os.MkdirAll(filepath.Join(repoPath, ".git"), 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}

	store := config.NewTestStore(root)
	if err := store.Save(config.Config{SearchPath: searchPath}); err != nil {
		t.Fatalf("save config: %v", err)
	}

	sel := &fakeSelector{choices: []selector.Choice{{Label: "alpha", Value: repoPath}}, ok: true}
	launcher := &fakeTmuxCommander{}
	deps := Dependencies{
		ConfigStore: store,
		Discoverer:  &discovery.Service{},
		Selector:    sel,
		Inspector: worktree.NewInspector(func(string) ([]byte, error) {
			return []byte("worktree " + repoPath + "\nbranch refs/heads/main\n\n"), nil
		}),
		TmuxLauncher: tmux.NewLauncher(launcher, nil, &bytes.Buffer{}, &bytes.Buffer{}),
		IsTTY:        func() bool { return true },
		Stdout:       &bytes.Buffer{},
		Stderr:       &bytes.Buffer{},
	}

	if err := NewRootCmd(deps).Execute(); err != nil {
		t.Fatalf("execute: %v", err)
	}
	if len(sel.titles) != 1 {
		t.Fatalf("expected one selector invocation, got %d", len(sel.titles))
	}
	if got := launcher.last[0]; got != repoPath {
		t.Fatalf("expected tmux dir %q, got %q", repoPath, got)
	}
}

func TestRootCommandShowsSecondSelectorForMultipleTargets(t *testing.T) {
	root := t.TempDir()
	searchPath := filepath.Join(root, "repos")
	repoPath := filepath.Join(searchPath, "alpha")
	worktreePath := filepath.Join(root, "alpha-feature")
	if err := os.MkdirAll(filepath.Join(repoPath, ".git"), 0o755); err != nil {
		t.Fatalf("mkdir repo: %v", err)
	}
	if err := os.MkdirAll(worktreePath, 0o755); err != nil {
		t.Fatalf("mkdir worktree: %v", err)
	}

	store := config.NewTestStore(root)
	if err := store.Save(config.Config{SearchPath: searchPath}); err != nil {
		t.Fatalf("save config: %v", err)
	}

	sel := &fakeSelector{choices: []selector.Choice{{Label: "alpha", Value: repoPath}, {Label: "alpha-feature", Value: worktreePath}}, ok: true}
	launcher := &fakeTmuxCommander{}
	deps := Dependencies{
		ConfigStore: store,
		Discoverer:  &discovery.Service{},
		Selector:    sel,
		Inspector: worktree.NewInspector(func(string) ([]byte, error) {
			return []byte("worktree " + repoPath + "\nbranch refs/heads/main\n\nworktree " + worktreePath + "\nbranch refs/heads/feature\n"), nil
		}),
		TmuxLauncher: tmux.NewLauncher(launcher, nil, &bytes.Buffer{}, &bytes.Buffer{}),
		IsTTY:        func() bool { return true },
		Stdout:       &bytes.Buffer{},
		Stderr:       &bytes.Buffer{},
	}

	if err := NewRootCmd(deps).Execute(); err != nil {
		t.Fatalf("execute: %v", err)
	}
	if len(sel.titles) != 2 {
		t.Fatalf("expected two selector invocations, got %d", len(sel.titles))
	}
	if got := launcher.last[0]; got != worktreePath {
		t.Fatalf("expected tmux dir %q, got %q", worktreePath, got)
	}
}

func TestRootCommandCancelExitsGracefully(t *testing.T) {
	root := t.TempDir()
	searchPath := filepath.Join(root, "repos")
	repoPath := filepath.Join(searchPath, "alpha")
	if err := os.MkdirAll(filepath.Join(repoPath, ".git"), 0o755); err != nil {
		t.Fatalf("mkdir: %v", err)
	}

	store := config.NewTestStore(root)
	if err := store.Save(config.Config{SearchPath: searchPath}); err != nil {
		t.Fatalf("save config: %v", err)
	}

	launcher := &fakeTmuxCommander{}
	deps := Dependencies{
		ConfigStore:  store,
		Discoverer:   &discovery.Service{},
		Selector:     &fakeSelector{ok: false, choices: []selector.Choice{{Label: "alpha", Value: repoPath}}},
		Inspector:    worktree.NewInspector(func(string) ([]byte, error) { return nil, nil }),
		TmuxLauncher: tmux.NewLauncher(launcher, nil, &bytes.Buffer{}, &bytes.Buffer{}),
		IsTTY:        func() bool { return true },
		Stdout:       &bytes.Buffer{},
		Stderr:       &bytes.Buffer{},
	}

	if err := NewRootCmd(deps).Execute(); err != nil {
		t.Fatalf("execute: %v", err)
	}
	if len(launcher.last) != 0 {
		t.Fatalf("expected no tmux invocation on cancel")
	}
}
