package commands

import (
	"fmt"
	"path/filepath"

	"github.com/ruggi/carmack/carmack"
	"github.com/urfave/cli"
)

// List shows a list of all the plan files.
func List(ctx *carmack.Context) cli.ActionFunc {
	return func(c *cli.Context) error {
		files, err := filepath.Glob(filepath.Join(ctx.UserFolder(), "*.plan"))
		if err != nil {
			return err
		}
		for _, f := range files {
			fmt.Println(f)
		}
		return nil
	}
}
