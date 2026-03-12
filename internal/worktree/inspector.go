package worktree

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

type Target struct {
	Name string
	Path string
}

type GitRunner func(repoPath string) ([]byte, error)

type Inspector struct {
	run GitRunner
}

func NewInspector(run GitRunner) *Inspector {
	return &Inspector{run: run}
}

func ExecGitRunner(repoPath string) ([]byte, error) {
	cmd := exec.Command("git", "-C", repoPath, "worktree", "list", "--porcelain")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return nil, fmt.Errorf("read git worktree metadata: %w: %s", err, strings.TrimSpace(string(output)))
	}
	return output, nil
}

func (i *Inspector) Targets(repoPath string) ([]Target, error) {
	output, err := i.run(repoPath)
	if err != nil {
		return nil, err
	}

	targets := parseTargets(output)
	if len(targets) == 0 {
		return nil, fmt.Errorf("read git worktree metadata: no worktree entries found")
	}

	return targets, nil
}

func parseTargets(output []byte) []Target {
	blocks := bytes.Split(bytes.TrimSpace(output), []byte("\n\n"))
	targets := make([]Target, 0, len(blocks))

	for index, block := range blocks {
		if len(bytes.TrimSpace(block)) == 0 {
			continue
		}

		var path string
		for _, line := range strings.Split(string(block), "\n") {
			if value, ok := strings.CutPrefix(line, "worktree "); ok {
				path = strings.TrimSpace(value)
				break
			}
		}
		if path == "" {
			continue
		}

		name := filepath.Base(path)
		if index == 0 {
			name += " (main)"
		}

		targets = append(targets, Target{Name: name, Path: path})
	}

	return targets
}
