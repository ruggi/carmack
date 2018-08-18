package commands

import (
	"fmt"
	"os"

	"github.com/ruggi/carmack/carmack"
	"github.com/ruggi/carmack/shell"
)

// Git runs a git command on the plan repo, if it's been initialized.
func Git(ctx *carmack.Context, args ...string) error {
	if !shell.Git.Initialized(ctx.Folder) {
		return fmt.Errorf("git not initialized, run '%s git init'", os.Args[0])
	}
	if len(args) == 0 {
		return fmt.Errorf("Usage: %s git [...]", os.Args[0])
	}
	return shell.Verbose.Git(ctx.Folder, args...)
}

// GitInit initializes the plan repo, if it's not already initialized.
func GitInit(ctx *carmack.Context) error {
	if shell.Git.Initialized(ctx.Folder) {
		return fmt.Errorf("git repo already initialized")
	}
	return shell.Git.Init(ctx.Folder)
}
