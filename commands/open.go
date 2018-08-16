package commands

import (
	"github.com/ruggi/carmack/shell"
	"github.com/urfave/cli"
)

// Open shows the list of open/pending entries in the plan files.
func Open(userFolder string) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		return shell.Verbose.Grep(userFolder, "^[^*+-]")
	}
}
