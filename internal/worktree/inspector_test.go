package worktree

import (
	"errors"
	"testing"
)

func TestInspectorParsesMainAndLinkedWorktrees(t *testing.T) {
	inspector := NewInspector(func(string) ([]byte, error) {
		return []byte("worktree /tmp/repos/alpha\nbranch refs/heads/main\n\nworktree /tmp/repos/alpha-feature\nbranch refs/heads/feature\n"), nil
	})

	targets, err := inspector.Targets("/tmp/repos/alpha")
	if err != nil {
		t.Fatalf("targets: %v", err)
	}
	if len(targets) != 2 {
		t.Fatalf("expected 2 targets, got %d", len(targets))
	}
	if targets[0].Name != "alpha (main)" || targets[1].Name != "alpha-feature" {
		t.Fatalf("unexpected targets: %+v", targets)
	}
}

func TestInspectorReturnsMetadataError(t *testing.T) {
	inspector := NewInspector(func(string) ([]byte, error) {
		return nil, errors.New("boom")
	})

	if _, err := inspector.Targets("/tmp/repos/alpha"); err == nil {
		t.Fatalf("expected error")
	}
}
