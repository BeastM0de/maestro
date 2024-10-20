package commands

import (
	"github.com/spf13/cobra"
	"log/slog"
	"maestro/internal/maestro/engine"
	"maestro/internal/maestro/git/commands"
)

func init() {
	rootCmd.AddCommand(pullCmd)
}

// pullCmd represents the command for pulling the latest changes from remote repositories.
// It executes 'git pull' for each project in the workspace that passes the filter criteria.
var pullCmd = &cobra.Command{
	Use:     "pull",
	Short:   "Pulls the latest changes from the remote repositories",
	Example: "maestro pull",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		slog.Info("Hello from pull command")

		filter := nil
		command := engine.Command{
			Cmd:    git.Pull,
			Filter: filter,
		}

		err := engine.Run(command)
		if err != nil {
			return err
		}

		return nil
	},
}
