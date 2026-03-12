package selector

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"

	"golang.org/x/term"

	"tm/internal/discovery"
	"tm/internal/worktree"
)

type Choice struct {
	Label   string
	Details string
	Value   string
}

func (c Choice) Title() string {
	return c.Label
}

func (c Choice) Description() string {
	return c.Details
}

func (c Choice) FilterValue() string {
	return strings.TrimSpace(c.Label + "\n" + c.Details)
}

type FzfSelector struct {
	in  *os.File
	out io.Writer
}

func NewFzfSelector(in *os.File, out io.Writer) *FzfSelector {
	return &FzfSelector{in: in, out: out}
}

func IsTTY() bool {
	stdinOK := term.IsTerminal(int(os.Stdin.Fd()))
	stdoutOK := term.IsTerminal(int(os.Stdout.Fd()))
	return stdinOK && stdoutOK
}

func RepoChoices(repos []discovery.Repo) []Choice {
	choices := make([]Choice, 0, len(repos))
	for _, repo := range repos {
		choices = append(choices, Choice{Label: repo.Name, Details: repo.Path, Value: repo.Path})
	}
	return choices
}

func TargetChoices(targets []worktree.Target) []Choice {
	choices := make([]Choice, 0, len(targets))
	for _, target := range targets {
		choices = append(choices, Choice{Label: target.Name, Details: target.Path, Value: target.Path})
	}
	return choices
}

func checkFzfInstalled() error {
	_, err := exec.LookPath("fzf")
	if err != nil {
		return fmt.Errorf("fzf is required but not installed. Please install fzf: https://github.com/junegunn/fzf#installation")
	}
	return nil
}

func (s *FzfSelector) Select(title string, items []Choice) (Choice, bool, error) {
	if err := checkFzfInstalled(); err != nil {
		return Choice{}, false, err
	}

	if len(items) == 0 {
		return Choice{}, false, nil
	}

	// Build input data: format is "Label\tDetails\tValue"
	var input bytes.Buffer
	for _, item := range items {
		// Display format: "Label: Details" for fzf
		display := item.Label
		if item.Details != "" {
			display = fmt.Sprintf("%s: %s", item.Label, item.Details)
		}
		// Store mapping: display text -> Choice
		// We use tab as delimiter between display and value
		fmt.Fprintf(&input, "%s\t%s\n", display, item.Value)
	}

	// Build fzf command
	args := []string{
		"--height=50%",
		"--reverse",
		"--border",
		"--prompt", title + "> ",
		"--delimiter=\t",
		"--with-nth=1", // Only show the display part, hide the value
		"--select-1",   // Auto-select if only one match
		"--exit-0",     // Exit 0 even if no match (we handle empty)
	}

	cmd := exec.Command("fzf", args...)
	cmd.Stdin = &input
	cmd.Stderr = s.out
	// Don't set cmd.Stdout - let Output() capture it

	output, err := cmd.Output()

	// Handle exit codes
	if exitErr, ok := err.(*exec.ExitError); ok {
		switch exitErr.ExitCode() {
		case 1:
			// No match
			return Choice{}, false, nil
		case 130:
			// Interrupted (Ctrl-C, Esc)
			return Choice{}, false, nil
		default:
			return Choice{}, false, fmt.Errorf("fzf failed with exit code %d", exitErr.ExitCode())
		}
	}

	if err != nil && !errors.Is(err, exec.ErrNotFound) {
		return Choice{}, false, fmt.Errorf("run fzf: %w", err)
	}

	// Parse output
	outputStr := strings.TrimSpace(string(output))
	if outputStr == "" {
		return Choice{}, false, nil
	}

	// Parse "display\tvalue" format
	parts := strings.SplitN(outputStr, "\t", 2)
	if len(parts) != 2 {
		return Choice{}, false, fmt.Errorf("unexpected fzf output format: %q", outputStr)
	}

	selectedValue := parts[1]

	// Find the matching choice
	for _, item := range items {
		if item.Value == selectedValue {
			return item, true, nil
		}
	}

	return Choice{}, false, fmt.Errorf("selected value not found in choices: %q", selectedValue)
}
