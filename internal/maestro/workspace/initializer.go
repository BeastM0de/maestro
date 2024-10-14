package workspace

import (
	"fmt"
	"github.com/spf13/viper"
	"log/slog"
	"maestro/internal/maestro/utils"
	"os"
	"path/filepath"
)

// InitializeWorkspace initializes a new workspace in either the current directory or a new directory with the name
// provided.
func InitializeWorkspace(name string) error {

	var workspaceHome, wErr = getWorkspaceHome(name)

	if wErr != nil {
		return wErr
	}
	slog.Debug("Workspace home: ", workspaceHome)

	//write the workspace configuration file
	cErr := createWorkspace(workspaceHome)

	return cErr
}

func getWorkspaceHome(name string) (string, error) {
	var workspaceHome string
	var currentWorkingDirectory, err = os.Getwd()
	if err != nil {
		return "", fmt.Errorf("error getting current working directory: " + err.Error())
	}

	if name != "" {
		workspaceHome = filepath.Join(currentWorkingDirectory, name)
	} else {
		workspaceHome = currentWorkingDirectory
	}
	return workspaceHome, nil
}

func createWorkspace(workspaceHome string) error {
	// Create the workspace directories
	err := utils.EnsurePathExists(workspaceHome)
	if err != nil {
		return fmt.Errorf("error creating src directory: " + err.Error())
	}
	// Specify the file name and type
	configFile := ".maestro"
	viper.SetConfigFile(configFile)
	viper.SetConfigType("yaml")

	// Write the configuration to the file
	err1 := viper.WriteConfigAs(configFile)
	if err1 != nil {
		return fmt.Errorf("error creating workpace: " + err1.Error())
	} else {
		fmt.Printf("Config file written successfully: %s\n", configFile)
	}

	return nil
}
