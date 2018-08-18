package git

import (
	"strings"

	"github.com/ruggi/carmack/shell"
)

// Initialized returns whether or not the git repo has been initialized.
func Initialized(folder string) bool {
	err := shell.Quiet.Git(folder, "rev-parse", "--is-inside-work-tree")
	return err == nil
}

// Init initializes the git repo.
func Init(folder string) error {
	err := shell.Verbose.Git(folder, "init")
	if err != nil {
		return err
	}
	err = Add(folder, ".")
	if err != nil {
		return err
	}
	err = Commit(folder, "initial commit")
	if err != nil {
		return err
	}
	return nil
}

// UserName returns the git user.name config value.
func UserName(folder string) string {
	b := shell.NewBuffered()
	err := b.Git(folder, "config", "user.name")
	if err != nil {
		return ""
	}
	return strings.TrimSpace(b.Buffer.String())
}

// Add adds files to the current stage.
func Add(folder string, files ...string) error {
	return shell.Verbose.Git(folder, append([]string{"add"}, files...)...)
}

// Commit commits the current stage.
func Commit(folder, message string) error {
	return shell.Verbose.Git(folder, "commit", "-m", message)
}
