package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:     "maestro",
	Version: "0.1.0",
	Short:   "Orchestrate your workspace like a maestro",
}

var pullCmd = &cobra.Command{
	Use:     "pull",
	Short:   "Pulls the latest changes from the remote repositories",
	Example: "maestro pull",
	Args: func(cmd *cobra.Command, args []string) error {
		return nil
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("Hello from get command")

		return nil
	},
}

func main() {
	rootCmd.AddCommand(pullCmd)

	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}
