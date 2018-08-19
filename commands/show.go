package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/ruggi/carmack/context"
	"github.com/ruggi/carmack/plan"
	"github.com/ruggi/carmack/shell"
)

// Show shows entries from plan files.
func Show(ctx *context.Context, user string, entryType plan.EntryType) error {
	targetFolder := ctx.UserFolder()
	if user != "" {
		targetFolder = filepath.Join(ctx.Folder, user)
		if _, err := os.Stat(targetFolder); os.IsNotExist(err) {
			return fmt.Errorf("user %q not found", user)
		}
	}

	var re string
	switch entryType {
	case plan.Done:
		re = "^\\*"
	case plan.Completed:
		re = "^\\+"
	case plan.Canceled:
		re = "^\\-"
	case plan.Note:
		re = "^[^*+-]"
	}
	if re == "" {
		return fmt.Errorf("missing re")
	}

	return shell.Verbose.Grep(targetFolder, re)
}
