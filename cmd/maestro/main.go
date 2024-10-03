package main

import (
	"maestro/internal/maestro/commands"
	"os"
)

func main() {

	err := commands.Execute()

	if err != nil {
		os.Exit(1)
	}
}
