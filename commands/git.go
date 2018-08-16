package commands

import (
	"fmt"
	"os"

	"github.com/ruggi/carmack/shell"
	"github.com/urfave/cli"
)

// Git runs a git command on the plan repo, if it's been initialized.
func Git(folder string) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		if !shell.HasGit(folder) {
			return fmt.Errorf("git not initialized, run '%s git init'", os.Args[0])
		}
		if len(ctx.Args()) == 0 {
			return fmt.Errorf("Usage: %s git [...]", os.Args[0])
		}
		return shell.Verbose.Git(folder, ctx.Args()...)
	}
}

// GitInit initializes the plan repo, if it's not already initialized.
func GitInit(folder string) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		if shell.HasGit(folder) {
			return fmt.Errorf("git repo already initialized")
		}
		err := shell.Verbose.Git(folder, "init")
		if err != nil {
			return err
		}
		err = shell.Verbose.Git(folder, "add", ".")
		if err != nil {
			return err
		}
		err = shell.Verbose.Git(folder, "commit", "-m", "initial commit")
		if err != nil {
			return err
		}
		return nil
	}
}
