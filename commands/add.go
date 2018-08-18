package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"github.com/ruggi/carmack/carmack"
	"github.com/ruggi/carmack/plan"
	"github.com/ruggi/carmack/shell"
)

const (
	timeFormat = "2006-01-02"
)

// Add adds a new entry to today's plan file.
func Add(ctx *carmack.Context, entry string, entryType plan.EntryType) error {
	if entry == "" {
		return fmt.Errorf("missing argument")
	}

	filename := filepath.Join(ctx.UserFolder(), time.Now().UTC().Format(timeFormat)+".plan")

	p, err := plan.Load(filename)
	if err != nil {
		return err
	}
	p.Add(entry, entryType)

	err = ioutil.WriteFile(filename, []byte(p.String()), os.ModePerm)
	if err != nil {
		return err
	}

	if shell.Git.Initialized(ctx.Folder) {
		err = shell.Git.Add(ctx.Folder, ".")
		if err != nil {
			return fmt.Errorf("cannot add: %s", err)
		}

		m := fmt.Sprintf(`'%s: plan update %s'`, ctx.Username, time.Now().UTC().Format(time.RFC3339))
		err = shell.Git.Commit(ctx.Folder, m)
		if err != nil {
			return fmt.Errorf("cannot commit: %s", err)
		}
	}

	return nil
}
