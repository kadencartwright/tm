package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func newConfigCmd(deps Dependencies) *cobra.Command {
	configCmd := &cobra.Command{
		Use:   "config",
		Short: "Manage tm configuration",
	}

	setCmd := &cobra.Command{
		Use:   "set",
		Short: "Set configuration values",
	}

	searchPathCmd := &cobra.Command{
		Use:   "search-path <path>",
		Short: "Set the repository discovery root",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			path, err := deps.ConfigStore.NormalizeSearchPath(args[0])
			if err != nil {
				return err
			}

			cfg, err := deps.ConfigStore.Load()
			if err != nil {
				return err
			}

			cfg.SearchPath = path
			if err := deps.ConfigStore.Save(cfg); err != nil {
				return err
			}

			_, _ = fmt.Fprintf(cmd.OutOrStdout(), "search_path set to %s\n", path)
			return nil
		},
	}

	searchPathCmd.ValidArgsFunction = func(*cobra.Command, []string, string) ([]string, cobra.ShellCompDirective) {
		return nil, cobra.ShellCompDirectiveDefault
	}

	setCmd.AddCommand(searchPathCmd)
	configCmd.AddCommand(setCmd)
	configCmd.SetOut(deps.Stdout)
	configCmd.SetErr(deps.Stderr)
	configCmd.SetIn(os.Stdin)

	return configCmd
}
