package commands

import (
	"github.com/spf13/cobra"
	"log/slog"
	"maestro/internal/maestro/workspace"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Adds the specified repository or group to the workspace",
	Example: "maestro add [name]",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		repo, err := workspace.AddResource(args[0])

		if err != nil {
			return err
		}
		slog.Info("Repository added successfully: ", repo.Name)
		return nil
	},
}
