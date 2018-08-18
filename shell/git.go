package shell

import (
	"strings"
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
func (g git) Init(folder string) error {
	err := Verbose.Git(folder, "init")
	if err != nil {
		return err
	}
	err = g.Add(folder, ".")
	if err != nil {
		return err
	}
	err = g.Commit(folder, "initial commit")
	if err != nil {
		return err
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
