package tmux

import (
	"errors"
	"io"
	"testing"
)

type fakeCommander struct {
	runArgs []string
	runDir  string
	lookErr error
	runErr  error
}

func (f *fakeCommander) LookPath(string) (string, error) {
	if f.lookErr != nil {
		return "", f.lookErr
	}
	return "/usr/bin/tmux", nil
}

func (f *fakeCommander) Run(args []string, dir string, _ io.Reader, _, _ io.Writer) error {
	f.runArgs = append([]string(nil), args...)
	f.runDir = dir
	return f.runErr
}

func TestSessionNameIsDeterministicAndDistinct(t *testing.T) {
	first := SessionName("/tmp/repos/alpha")
	second := SessionName("/tmp/repos/alpha")
	third := SessionName("/tmp/repos/alpha-feature")

	if first != second {
		t.Fatalf("expected deterministic session names")
	}
	if first == third {
		t.Fatalf("expected distinct session names")
	}
}

func TestAttachOrCreateReportsMissingTmux(t *testing.T) {
	launcher := NewLauncher(&fakeCommander{lookErr: errors.New("missing")}, nil, nil, nil)
	if err := launcher.AttachOrCreate("/tmp/repos/alpha"); err == nil {
		t.Fatalf("expected error")
	}
}

func TestAttachOrCreateReportsTmuxFailure(t *testing.T) {
	launcher := NewLauncher(&fakeCommander{runErr: errors.New("boom")}, nil, nil, nil)
	if err := launcher.AttachOrCreate("/tmp/repos/alpha"); err == nil {
		t.Fatalf("expected error")
	}
}
