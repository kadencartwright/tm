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
	DetachClient() error
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

func (ExecCommander) DetachClient() error {
	cmd := exec.Command("tmux", "detach-client")
	return cmd.Run()
}

// ReExecute re-runs the current tm command from outside the tmux session context.
// It should be called after detaching from the current tmux session.
func ReExecute() error {
	execPath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to get executable path: %w", err)
	}

	args := os.Args
	cmd := exec.Command(execPath, args[1:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

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

func (l *Launcher) AttachOrCreate(targetPath string) error {
	if _, err := l.commander.LookPath("tmux"); err != nil {
		return fmt.Errorf("tmux is required but not available: %w", err)
	}

	// Check if running inside a tmux session and stdin is a TTY
	// This avoids triggering detach in non-interactive scripts
	if IsNestedSession() && IsTerminal() {
		// Detach from current session
		if err := l.commander.DetachClient(); err != nil {
			return fmt.Errorf("failed to detach from current tmux session: %w", err)
		}
		// Re-execute the tm command outside the tmux context
		if err := ReExecute(); err != nil {
			return fmt.Errorf("failed to re-execute command: %w", err)
		}
		// ReExecute replaces the process, so we won't reach here on success
		return nil
	}

	session := SessionName(targetPath)
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
