package commands

import (
	"fmt"
	"path/filepath"

	"github.com/urfave/cli"
)

// List shows a list of all the plan files.
func List(userFolder string) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		files, err := filepath.Glob(filepath.Join(userFolder, "*.plan"))
		if err != nil {
			return err
		}
		for _, f := range files {
			fmt.Println(f)
		}
		return nil
	}
}
