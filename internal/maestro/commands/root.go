package commands

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log/slog"
	"os"
)

var (
	// Used for flags.
	cfgFile string
	Verbose bool

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

	rootCmd.PersistentFlags().StringVar(&cfgFile, "dry-run", "", "Run in dry-run mode where changes or commands are not applied")
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.maestro.yaml)")
	rootCmd.PersistentFlags().Bool("viper", true, "use Viper for configuration")
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
	viper.BindPFlag("useViper", rootCmd.PersistentFlags().Lookup("viper"))
	viper.Set("name", "maestro")
	viper.Set("version", "0.0.1")

	if Verbose == true {
		slog.SetLogLoggerLevel(slog.LevelDebug)
	}

}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		slog.Debug("Config file set: ", cfgFile)
		viper.SetConfigFile(cfgFile)
	} else {
		slog.Debug("No config file set, using default $HOME/.maestro.yaml")
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
		slog.Debug("Using config file:", viper.ConfigFileUsed())
	}
}
