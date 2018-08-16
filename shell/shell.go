package shell

import (
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

// Verbose is a verbose (stdin, stdout, stderr) command issuer.
var Verbose = C{
	Stdin:  os.Stdin,
	Stdout: os.Stdout,
	Stderr: os.Stderr,
}

// Quiet is a quiet (no output) command issuer.
var Quiet = C{
	Stdin: os.Stdin,
}

// C is a command issuer.
type C struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

func (c C) run(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdin = c.Stdin
	cmd.Stdout = c.Stdout
	cmd.Stderr = c.Stderr
	return cmd.Run()
}

// Git runs git commands on the given folder.
func (c C) Git(folder string, args ...string) error {
	return c.run("git", append([]string{"-C", folder}, args...)...)
}

// Grep runs grep on all the plan files in the given folder.
func (c C) Grep(userFolder string, re string) error {
	files, err := filepath.Glob(filepath.Join(userFolder, "*.plan"))
	if err != nil {
		return err
	}
	if len(files) == 0 {
		return nil
	}
	return c.run("grep", append([]string{re}, files...)...)
}

// HasGit returns whether the plan repo has been initialized or not.
func HasGit(folder string) bool {
	err := Quiet.Git(folder, "rev-parse", "--is-inside-work-tree")
	return err == nil
}
