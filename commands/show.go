package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ruggi/carmack/shell"
	"github.com/urfave/cli"
)

// Show shows entries in plan files.
func Show(folder, userFolder string) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		targetFolder := userFolder
		if ctx.IsSet("user") {
			u := ctx.String("user")
			if _, err := os.Stat(targetFolder); os.IsNotExist(err) {
				return fmt.Errorf("user %q not found", u)
			}
			targetFolder = filepath.Join(folder, u)
		}

		var re string
		if ctx.Bool("open") {
			re = "^[^*+-]"
		} else if ctx.Bool("done") {
			re = "^\\*"
		} else if ctx.Bool("completed") {
			re = "^\\+"
		} else if ctx.Bool("canceled") {
			re = "^\\-"
		}
		if re == "" {
			return fmt.Errorf("missing re")
		}

		return shell.Verbose.Grep(targetFolder, re)
	}
}
