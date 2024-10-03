package commands

import (
	"github.com/spf13/cobra"
	"log/slog"
)

func init() {
	initCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.maestro.yaml)")
	initCmd.PersistentFlags().StringVar(&interactive, "interactive", "i", "", "interactive mode")
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:     "init",
	Short:   "Initialize a new workspace, either from scratch or from a configuration file",
	Example: "maestro init",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		slog.Debug("Hello from init command, using config file: ", cfgFile)

		return nil
	},
}
