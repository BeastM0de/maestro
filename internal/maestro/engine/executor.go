package engine

import "maestro/internal/maestro"

type Command func(project maestro.Repository) error
type Filter func(project string) bool

func run(cmd Command, filter Filter) error {
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
