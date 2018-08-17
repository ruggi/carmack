package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ruggi/carmack/config"
	"github.com/ruggi/carmack/git"
	"github.com/ruggi/carmack/plan"
	"github.com/urfave/cli"
)

const (
	timeFormat = "2006-01-02"
)

type Carmack struct {
	conf config.Config
}

// Add inserts a new entry to today's plan file.
func Add(ctx *cli.Context) error {
		if len(ctx.Args()) == 0 {
			return fmt.Errorf("missing argument")
		}
		s := strings.Join(ctx.Args(), " ")
		filename := filepath.Join(userFolder, time.Now().UTC().Format(timeFormat)+".plan")
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
		if git.Initialized(folder) {
			err = git.Add(folder, ".")
			if err != nil {
				return fmt.Errorf("cannot add: %s", err)
			}
			err = git.Commit(folder, fmt.Sprintf(`'%s: plan update %s'`, username, time.Now().UTC().Format(time.RFC3339)))
			if err != nil {
				return fmt.Errorf("cannot commit: %s", err)
			}
		}
		return nil
	}
}
