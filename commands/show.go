package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ruggi/carmack/carmack"
	"github.com/ruggi/carmack/shell"
	"github.com/urfave/cli"
)

// Show shows entries in plan files.
func Show(ctx *carmack.Context) cli.ActionFunc {
	return func(c *cli.Context) error {
		targetFolder := ctx.UserFolder()
		if c.IsSet("user") {
			u := c.String("user")
			if _, err := os.Stat(targetFolder); os.IsNotExist(err) {
				return fmt.Errorf("user %q not found", u)
			}
			targetFolder = filepath.Join(ctx.Folder, u)
		}

		var re string
		if c.Bool("open") {
			re = "^[^*+-]"
		} else if c.Bool("done") {
			re = "^\\*"
		} else if c.Bool("completed") {
			re = "^\\+"
		} else if c.Bool("canceled") {
			re = "^\\-"
		}
		if re == "" {
			return fmt.Errorf("missing re")
		}

		return shell.Verbose.Grep(targetFolder, re)
	}
}
