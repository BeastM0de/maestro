package engine

import "maestro/internal/maestro"

type Cmd func(project maestro.Repository) error
type Filter func(project string) bool

type Command struct {
	Cmd Cmd
	Filter Filter
}

func GitPullCommand() Command {
	return func(project maestro.Repository) error {
		return project.Pull()
	}
}

func Run(cmd Command) error {
	for project := range projects {
		if filter != nil {
			if !filter(project) {
				continue
			}
		}
		if err := cmd(project); err != nil {
			return err
		}

	return nil
}