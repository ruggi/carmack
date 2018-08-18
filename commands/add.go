package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ruggi/carmack/carmack"
	"github.com/ruggi/carmack/git"
	"github.com/ruggi/carmack/plan"
	"github.com/urfave/cli"
)

const (
	timeFormat = "2006-01-02"
)

// Add inserts a new entry to today's plan file.
func Add(ctx *carmack.Context) cli.ActionFunc {
	return func(c *cli.Context) error {
		if len(c.Args()) == 0 {
			return fmt.Errorf("missing argument")
		}

		filename := filepath.Join(ctx.UserFolder(), time.Now().UTC().Format(timeFormat)+".plan")
		p, err := plan.Load(filename)
		if err != nil {
			return err
		}

		s := strings.Join(c.Args(), " ")
		if c.Bool("done") {
			p.AddDone(s)
		} else if c.Bool("completed") {
			p.AddCompleted(s)
		} else if c.Bool("canceled") {
			p.AddCanceled(s)
		} else {
			p.AddNote(s)
		}

		err = ioutil.WriteFile(filename, []byte(p.String()), os.ModePerm)
		if err != nil {
			return err
		}

		if git.Initialized(ctx.Folder) {
			err = git.Add(ctx.Folder, ".")
			if err != nil {
				return fmt.Errorf("cannot add: %s", err)
			}

			m := fmt.Sprintf(`'%s: plan update %s'`, ctx.Username, time.Now().UTC().Format(time.RFC3339))
			err = git.Commit(ctx.Folder, m)
			if err != nil {
				return fmt.Errorf("cannot commit: %s", err)
			}
		}

		return nil
	}
}
