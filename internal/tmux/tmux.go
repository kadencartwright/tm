package tmux

import (
	"crypto/sha1"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"

	"golang.org/x/term"
)

type Commander interface {
	LookPath(file string) (string, error)
	Run(args []string, dir string, in io.Reader, out, errOut io.Writer) error
}

type ExecCommander struct{}

func (ExecCommander) LookPath(file string) (string, error) {
	return exec.LookPath(file)
}

func (ExecCommander) Run(args []string, dir string, in io.Reader, out, errOut io.Writer) error {
	cmd := exec.Command("tmux", args...)
	cmd.Dir = dir
	cmd.Stdin = in
	cmd.Stdout = out
	cmd.Stderr = errOut
	return cmd.Run()
}

type Launcher struct {
	commander Commander
	in        io.Reader
	out       io.Writer
	errOut    io.Writer
}

func NewLauncher(commander Commander, in io.Reader, out, errOut io.Writer) *Launcher {
	return &Launcher{commander: commander, in: in, out: out, errOut: errOut}
}

func SessionName(path string) string {
	clean := filepath.Clean(path)
	base := sanitize(filepath.Base(clean))
	if base == "" {
		base = "tm"
	}
	hash := sha1.Sum([]byte(clean))
	return fmt.Sprintf("%s-%x", base, hash[:4])
}

// IsNestedSession returns true if the current process is running inside a tmux session
func IsNestedSession() bool {
	return os.Getenv("TMUX") != ""
}

// IsTerminal returns true if stdin is a TTY
func IsTerminal() bool {
	return term.IsTerminal(int(os.Stdin.Fd()))
}

// SessionExists checks if a tmux session with the given name already exists
func (l *Launcher) SessionExists(session string) bool {
	args := []string{"has-session", "-t", session}
	err := l.commander.Run(args, "", l.in, l.out, l.errOut)
	return err == nil
}

func (l *Launcher) AttachOrCreate(targetPath string) error {
	if _, err := l.commander.LookPath("tmux"); err != nil {
		return fmt.Errorf("tmux is required but not available: %w", err)
	}

	session := SessionName(targetPath)

	// Check if running inside a tmux session
	if IsNestedSession() && IsTerminal() {
		// When inside tmux, we need to:
		// 1. Create the session if it doesn't exist (detached)
		// 2. Switch to it using switch-client

		if !l.SessionExists(session) {
			// Create session detached (-d flag)
			createArgs := []string{"new-session", "-d", "-s", session, "-c", targetPath}
			if err := l.commander.Run(createArgs, targetPath, l.in, l.out, l.errOut); err != nil {
				return fmt.Errorf("tmux failed to create session %q: %w", session, err)
			}
		}

		// Switch to the session
		switchArgs := []string{"switch-client", "-t", session}
		if err := l.commander.Run(switchArgs, targetPath, l.in, l.out, l.errOut); err != nil {
			return fmt.Errorf("tmux failed to switch to session %q: %w", session, err)
		}

		return nil
	}

	// Not inside tmux - use standard attach/create
	args := []string{"new-session", "-A", "-s", session, "-c", targetPath}
	if err := l.commander.Run(args, targetPath, l.in, l.out, l.errOut); err != nil {
		return fmt.Errorf("tmux failed for session %q: %w", session, err)
	}

	return nil
}

var invalidSessionChars = regexp.MustCompile(`[^a-zA-Z0-9_-]+`)

func sanitize(value string) string {
	value = invalidSessionChars.ReplaceAllString(value, "-")
	return strings.Trim(value, "-")
}
