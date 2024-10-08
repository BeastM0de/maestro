package commands

import (
	"github.com/spf13/cobra"
	"log/slog"
)

func init() {
	rootCmd.AddCommand(configCmd)
}

var configCmd = &cobra.Command{
	Use:     "config",
	Short:   "View or modify the workspace configuration",
	Example: "maestro config",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		slog.Debug("Hello from config command, using config file: ", cfgFile)

		return nil
	},
}
