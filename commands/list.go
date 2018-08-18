package commands

import (
	"fmt"
	"path/filepath"

	"github.com/ruggi/carmack/carmack"
)

// List shows a list of all plan files.
func List(ctx *carmack.Context) error {
	files, err := filepath.Glob(filepath.Join(ctx.UserFolder(), "*.plan"))
	if err != nil {
		return err
	}
	for _, f := range files {
		fmt.Println(f)
	}
	return nil
}
