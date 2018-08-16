package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ruggi/carmack/plan"
	"github.com/ruggi/carmack/shell"
	"github.com/urfave/cli"
)

const (
	timeFormat = "2006-01-02"
)

// Add inserts a new entry to today's plan file.
func Add(username, folder, userFolder string) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		if len(ctx.Args()) == 0 {
			return fmt.Errorf("missing argument")
		}
		s := strings.Join(ctx.Args(), " ")
		filename := filepath.Join(userFolder, time.Now().Format(timeFormat)+".plan")
		p, err := plan.Load(filename)
		if err != nil {
			return err
		}
		if ctx.Bool("done") {
			p.AddDone(s)
		} else if ctx.Bool("completed") {
			p.AddCompleted(s)
		} else if ctx.Bool("canceled") {
			p.AddCanceled(s)
		} else {
			p.AddNote(s)
		}
		err = ioutil.WriteFile(filename, []byte(p.String()), os.ModePerm)
		if err != nil {
			return err
		}
		if shell.HasGit(folder) {
			err = shell.Verbose.Git(folder, "add", ".")
			if err != nil {
				return fmt.Errorf("cannot add: %s", err)
			}
			err = shell.Verbose.Git(folder, "commit", "-m", fmt.Sprintf(`'%s: plan update %s'`, username, time.Now().Format(time.RFC3339)))
			if err != nil {
				return fmt.Errorf("cannot commit: %s", err)
			}
		}
		return nil
	}
}
