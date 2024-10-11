package workspace

import (
	"log/slog"
	"maestro/internal/maestro/utils"
	"os"
	"path/filepath"
)

func InitializeWorkspace(name string) {

	var workspaceHome, err = createWorkspaceHome(name)

	if err != nil {
		slog.Error("Error creating workspace home: ", err)
		return
	}
	slog.Debug("Workspace home: ", workspaceHome)

}

func createWorkspaceHome(name string) (string, error) {
	var workspaceHome string
	var currentWorkingDirectory, err = os.Getwd()
	if err != nil {
		return "", Error("Error getting current working directory")
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
