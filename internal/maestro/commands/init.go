package commands

import (
	"github.com/spf13/cobra"
	"log/slog"
	"maestro/internal/maestro/workspace"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:     "init [name]",
	Short:   "Initialize a new workspace.",
	Long:    "Initialize a new workspace in either the current directory or a new directory with the name provided.",
	Example: "maestro init",
	Args:    cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		slog.Debug("Hello from init command, using config file: ", cfgFile)

		var name string
		if len(args) == 1 {
			name = args[0]
		}

		err := workspace.InitializeWorkspace(name)

		if err != nil {
			slog.Error("Error initializing workspace: ", err)
		}
		return err
	},
}
