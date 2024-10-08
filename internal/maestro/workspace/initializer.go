package workspace

import (
	"log/slog"
	"os"
)

func InitializeWorkspace(name string) {
	var currentWorkingDirectory string = getCurrentWorkingDirectory()
	var workspaceHome string

	if name == "" {
		// Initialize workspace in current directory
	} else {
		// Initialize workspace in new directory with the name provided
	}

}

func getCurrentWorkingDirectory() string {
	path, err := os.Getwd()
	if err != nil {
		slog.Println("Error getting current working directory")
		return ""
	} else {
	}
	return path
}
