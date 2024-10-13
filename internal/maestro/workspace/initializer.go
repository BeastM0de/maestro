package workspace

import (
	"fmt"
	"log/slog"
	"maestro/internal/maestro/utils"
	"os"
	"path/filepath"
)

func InitializeWorkspace(name string) error {

	var workspaceHome, err = createWorkspaceHome(name)

	if err != nil {
		return err
	}
	slog.Debug("Workspace home: ", workspaceHome)

	return err
}

func createWorkspaceHome(name string) (string, error) {
	var workspaceHome string
	var currentWorkingDirectory, err = os.Getwd()
	if err != nil {
		return "", fmt.Errorf("error getting current working directory")
	}

	if name != "" {
		workspaceHome = filepath.Join(currentWorkingDirectory, name)

		err := utils.EnsurePathExists(workspaceHome)
		if err != nil {
			return "", err
		}

	} else {
		workspaceHome = currentWorkingDirectory
	}
	return workspaceHome, nil
}
