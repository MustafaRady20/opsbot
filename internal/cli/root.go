package cli

import "github.com/spf13/cobra"

// NewRootCmd builds the full opsbot command tree. version is printed by
// the version command.
func NewRootCmd(version string) *cobra.Command {
	var configPath string

	root := &cobra.Command{
		Use:   "opsbot",
		Short: "Monitor servers and sites, and run a personal Telegram bot",
		// Don't print the usage text when a command fails at runtime,
		// only when the user typed the command wrong.
		SilenceUsage: true,
	}

	// Persistent: available to every subcommand, as in `opsbot check --config x.yaml`.
	// T04: pass configPath to commands that load the config.
	root.PersistentFlags().StringVar(&configPath, "config", "config.yaml", "path to the config file")

	root.AddCommand(
		newCheckCmd(),
		newVersionCmd(version),
	)
	return root
}
