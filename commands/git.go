package commands

import (
	"fmt"
	"os"

	"github.com/ruggi/carmack/carmack"
	"github.com/ruggi/carmack/git"
	"github.com/ruggi/carmack/shell"
	"github.com/urfave/cli"
)

// Git runs a git command on the plan repo, if it's been initialized.
func Git(ctx *carmack.Context) cli.ActionFunc {
	return func(c *cli.Context) error {
		if !git.Initialized(ctx.Folder) {
			return fmt.Errorf("git not initialized, run '%s git init'", os.Args[0])
		}
		if len(c.Args()) == 0 {
			return fmt.Errorf("Usage: %s git [...]", os.Args[0])
		}
		return shell.Verbose.Git(ctx.Folder, c.Args()...)
	}
}

// GitInit initializes the plan repo, if it's not already initialized.
func GitInit(ctx *carmack.Context) cli.ActionFunc {
	return func(c *cli.Context) error {
		if git.Initialized(ctx.Folder) {
			return fmt.Errorf("git repo already initialized")
		}
		return git.Init(ctx.Folder)
	}
}
