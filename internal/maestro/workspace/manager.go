package workspace

import (
	"log/slog"
	"maestro/internal/maestro"
)

func AddResource(name string) (maestro.Resource, error) {
	slog.Debug("Adding resource", name)
	// Add a repository or group to the workspace

	//determine resource type

	var resource maestro.Repository
	return nil, nil
}

func AddGroup(group maestro.Group) {
	// Add a group to the workspace
}
func AddRepository(repository maestro.Repository) {
	// Add a repository to the workspace
}
