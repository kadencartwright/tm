package tmux

import (
	"crypto/sha1"
	"fmt"
	"io"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
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

func (l *Launcher) AttachOrCreate(targetPath string) error {
	if _, err := l.commander.LookPath("tmux"); err != nil {
		return fmt.Errorf("tmux is required but not available: %w", err)
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
