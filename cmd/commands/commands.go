package commands

import (
	"github.com/urfave/cli/v3"
)

var commands []*cli.Command

func Commands() []*cli.Command {
	register(migrationsUpCommand)
	register(migrationsDownCommand)

	return commands
}

func register(command *cli.Command) {
	commands = append(commands, command)
}
