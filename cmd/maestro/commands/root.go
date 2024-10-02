package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	// Used for flags.
	cfgFile     string
	userLicense string
	verbose     *bool

	rootCmd = &cobra.Command{
		Use:   "maestro",
		Short: "An orchestration tool for your workspace",
		Long:  `Maestro is a CLI tool that helps you manage your workspace`,
	}
)

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.maestro.yaml)")
	rootCmd.PersistentFlags().StringP("author", "a", "Jon Cappella", "author name for copyright attribution")
	rootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "GPLv3", "name of license for the project")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	verbose = rootCmd.PersistentFlags().Bool("verbose", false, "verbose output")
	viper.BindPFlag("author", rootCmd.PersistentFlags().Lookup("author"))
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.SetDefault("author", "Jon Cappella <jonacappella@gmail.com>")
	viper.SetDefault("license", "GPLv3")

}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		logVerbose("Config file set: ", cfgFile)
		viper.SetConfigFile(cfgFile)
	} else {
		logVerbose("No config file set, using default $HOME/.maestro.yaml")
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cobra" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".maestro")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func logVerbose(a ...any) {
	if *verbose == true {
		fmt.Println(a...)
	}
}
