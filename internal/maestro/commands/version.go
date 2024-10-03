package commands

import (
	"github.com/spf13/cobra"
	"log/slog"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Maestro",
	Long:  `All software has versions. This is Maestro's`,
	Run: func(cmd *cobra.Command, args []string) {
		slog.Info("Maestro v0.1.0")
	},
}
