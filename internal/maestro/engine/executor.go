package engine

import (
	"log/slog"
	"maestro/internal/maestro"
)

type ProjectCommand func(project maestro.Repository) error
type Filter func(project string) bool

type Command struct {
	Run    ProjectCommand
	Filter Filter
}

func Run(command Command) error {
	for project := range projects {
		if command.Filter != nil {
			if !command.Filter(project) {
				slog.Debug("Skipping project %s", project.Name)
				continue
			}
		}
		if err := command.Run(project); err != nil {
			return err
		}
	}
	return nil
}
