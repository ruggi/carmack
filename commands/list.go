package commands

import (
	"fmt"
	"path/filepath"

	"github.com/ruggi/carmack/context"
)

// List shows a list of all plan files.
func List(ctx *context.Context) error {
	files, err := filepath.Glob(filepath.Join(ctx.UserFolder(), "*.plan"))
	if err != nil {
		return err
	}
	for _, f := range files {
		fmt.Println(f)
	}
	return nil
}
