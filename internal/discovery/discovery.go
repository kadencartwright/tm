package discovery

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Repo struct {
	Name string
	Path string
}

type Service struct{}

func (s *Service) Discover(searchPath string) ([]Repo, error) {
	entries, err := os.ReadDir(searchPath)
	if err != nil {
		return nil, fmt.Errorf("read search path: %w", err)
	}

	repos := make([]Repo, 0)
	for _, entry := range entries {
		if !entry.IsDir() || strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		repoPath := filepath.Join(searchPath, entry.Name())
		if _, err := os.Stat(filepath.Join(repoPath, ".git")); err != nil {
			if os.IsNotExist(err) {
				continue
			}
			return nil, fmt.Errorf("inspect repository metadata: %w", err)
		}

		repos = append(repos, Repo{Name: entry.Name(), Path: repoPath})
	}

	sort.Slice(repos, func(i, j int) bool {
		return repos[i].Name < repos[j].Name
	})

	return repos, nil
}
