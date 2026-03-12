package cmd

import (
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"

	"tm/internal/config"
	"tm/internal/discovery"
	"tm/internal/selector"
	"tm/internal/tmux"
	"tm/internal/worktree"
)

type Selector interface {
	Select(title string, items []selector.Choice) (selector.Choice, bool, error)
}

type Dependencies struct {
	ConfigStore  *config.Store
	Discoverer   *discovery.Service
	Selector     Selector
	Inspector    *worktree.Inspector
	TmuxLauncher *tmux.Launcher
	IsTTY        func() bool
	Stdout       io.Writer
	Stderr       io.Writer
}

func DefaultDependencies() Dependencies {
	return Dependencies{
		ConfigStore:  config.NewStore(),
		Discoverer:   &discovery.Service{},
		Selector:     selector.NewFzfSelector(os.Stdin, os.Stdout),
		Inspector:    worktree.NewInspector(worktree.ExecGitRunner),
		TmuxLauncher: tmux.NewLauncher(tmux.ExecCommander{}, os.Stdin, os.Stdout, os.Stderr),
		IsTTY:        selector.IsTTY,
		Stdout:       os.Stdout,
		Stderr:       os.Stderr,
	}
}

func NewRootCmd(deps Dependencies) *cobra.Command {
	cmd := &cobra.Command{
		Use:           "tm",
		Short:         "Jump from repository selection into tmux",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return runRoot(deps)
		},
	}

	cmd.SetOut(deps.Stdout)
	cmd.SetErr(deps.Stderr)
	cmd.AddCommand(newConfigCmd(deps))

	return cmd
}

func runRoot(deps Dependencies) error {
	cfg, err := deps.ConfigStore.Load()
	if err != nil {
		return err
	}

	if cfg.SearchPath == "" {
		return fmt.Errorf("missing search path configuration; run `tm config set search-path <path>`")
	}

	repos, err := deps.Discoverer.Discover(cfg.SearchPath)
	if err != nil {
		return err
	}
	if len(repos) == 0 {
		return fmt.Errorf("no repositories found under %q", cfg.SearchPath)
	}

	if !deps.IsTTY() {
		return fmt.Errorf("interactive selection requires a TTY")
	}

	repoChoices := selector.RepoChoices(repos)
	selectedRepo, ok, err := deps.Selector.Select("Select a repository", repoChoices)
	if err != nil {
		return err
	}
	if !ok {
		return nil
	}

	targets, err := deps.Inspector.Targets(selectedRepo.Value)
	if err != nil {
		return err
	}
	if len(targets) == 0 {
		return fmt.Errorf("no checkout targets found for %q", selectedRepo.Value)
	}

	selectedTarget := targets[0]
	if len(targets) > 1 {
		targetChoice, targetOK, selectErr := deps.Selector.Select("Select a checkout target", selector.TargetChoices(targets))
		if selectErr != nil {
			return selectErr
		}
		if !targetOK {
			return nil
		}
		selectedTarget = worktree.Target{Name: targetChoice.Label, Path: targetChoice.Value}
	}

	if err := deps.TmuxLauncher.AttachOrCreate(selectedTarget.Path); err != nil {
		return err
	}

	return nil
}
