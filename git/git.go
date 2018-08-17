package git

import (
	"strings"

	"github.com/ruggi/carmack/shell"
)

func Initialized(folder string) bool {
	err := shell.Quiet.Git(folder, "rev-parse", "--is-inside-work-tree")
	return err == nil
}

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

func UserName(folder string) string {
	b := shell.NewBuffered()
	err := b.Git(folder, "config", "user.name")
	if err != nil {
		return ""
	}
	return strings.TrimSpace(b.Buffer.String())
}

func Add(folder string, files ...string) error {
	return shell.Verbose.Git(folder, append([]string{"add"}, files...)...)
}

func Commit(folder, message string) error {
	return shell.Verbose.Git(folder, "commit", "-m", message)
}
