package discovery

import (
	"os"
	"path/filepath"
	"testing"
)

func TestDiscoverFiltersToRepositories(t *testing.T) {
	root := t.TempDir()
	for _, dir := range []string{"alpha", "beta", "plain", ".hidden"} {
		if err := os.MkdirAll(filepath.Join(root, dir), 0o755); err != nil {
			t.Fatalf("mkdir %s: %v", dir, err)
		}
	}
	if err := os.MkdirAll(filepath.Join(root, "alpha", ".git"), 0o755); err != nil {
		t.Fatalf("mkdir alpha git: %v", err)
	}
	if err := os.WriteFile(filepath.Join(root, "beta", ".git"), []byte("gitdir: /tmp/beta"), 0o644); err != nil {
		t.Fatalf("write beta git file: %v", err)
	}

	repos, err := (&Service{}).Discover(root)
	if err != nil {
		t.Fatalf("discover: %v", err)
	}
	if len(repos) != 2 {
		t.Fatalf("expected 2 repos, got %d", len(repos))
	}
	if repos[0].Name != "alpha" || repos[1].Name != "beta" {
		t.Fatalf("unexpected repos: %+v", repos)
	}
}
