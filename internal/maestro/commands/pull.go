package commands

import (
	"github.com/spf13/cobra"
	"log/slog"
)

func init() {
	rootCmd.AddCommand(pullCmd)
}

var pullCmd = &cobra.Command{
	Use:     "pull",
	Short:   "Pulls the latest changes from the remote repositories",
	Example: "maestro pull",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		slog.Info("Hello from pull command")

		return nil
	},
}
