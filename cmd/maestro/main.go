package main

import (
	"maestro/cmd/maestro/commands"
	"os"
)

func main() {

	err := commands.Execute()

	if err != nil {
		os.Exit(1)
	}
}
