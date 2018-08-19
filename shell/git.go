package shell

import (
	"path/filepath"
	"strings"

	"github.com/ruggi/carmack/context"
	"github.com/ruggi/carmack/util"
)

type git struct{}

// Git is a wrapper around common git commands.
var Git = git{}

// Initialized returns whether or not the git repo has been initialized.
func (g git) Initialized(folder string) bool {
	err := Quiet.Git(folder, "rev-parse", "--is-inside-work-tree")
	return err == nil
}

// Init initializes the git repo.
func (g git) Init(ctx *context.Context) error {
	err := Verbose.Git(ctx.Folder, "init")
	if err != nil {
		return err
	}
	files, err := filepath.Glob(filepath.Join(ctx.UserFolder(), "*.plan"))
	if err != nil {
		return err
	}
	if len(files) > 0 {
		err = g.Add(ctx.Folder, files...)
		if err != nil {
			return err
		}
		err = g.Commit(ctx.Folder, util.CommitMessage(ctx))
		if err != nil {
			return err
		}
	}
	return nil
}

// UserName returns the git user.name config value.
func (g git) UserName(folder string) string {
	b := NewBuffered()
	err := b.Git(folder, "config", "--local", "user.name")
	if err != nil {
		b.Buffer.Reset()
		err = b.Git(".", "config", "--global", "user.name")
		if err != nil {
			return ""
		}
	}
	return strings.TrimSpace(b.Buffer.String())
}

// Add adds files to the current stage.
func (g git) Add(folder string, files ...string) error {
	return Verbose.Git(folder, append([]string{"add"}, files...)...)
}

// Commit commits the current stage.
func (g git) Commit(folder, message string) error {
	return Verbose.Git(folder, "commit", "-m", message)
}
