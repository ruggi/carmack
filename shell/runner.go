package shell

import (
	"bytes"
	"io"
	"os"
	"os/exec"
	"path/filepath"
)

// Verbose is a verbose (stdin, stdout, stderr) runner.
var Verbose = Runner{
	Stdin:  os.Stdin,
	Stdout: os.Stdout,
	Stderr: os.Stderr,
}

// Quiet is a quiet (no output) runner.
var Quiet = Runner{
	Stdin: os.Stdin,
}

// Buffered is a runner which writes on a bytes.Buffer.
type Buffered struct {
	Runner
	Buffer bytes.Buffer
}

// NewBuffered creates a new buffered runner.
func NewBuffered() *Buffered {
	b := Buffered{}
	b.Runner.Stdin = os.Stdin
	b.Runner.Stdout = &b.Buffer
	b.Runner.Stderr = &b.Buffer
	return &b
}

// Runner is a shell command runner.
type Runner struct {
	Stdin  io.Reader
	Stdout io.Writer
	Stderr io.Writer
}

func (r Runner) run(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdin = r.Stdin
	cmd.Stdout = r.Stdout
	cmd.Stderr = r.Stderr
	return cmd.Run()
}

// Git runs git commands on the given folder.
func (r Runner) Git(folder string, args ...string) error {
	return r.run("git", append([]string{"-C", folder}, args...)...)
}

// Grep runs grep on all the plan files in the given folder.
func (r Runner) Grep(userFolder string, re string) error {
	files, err := filepath.Glob(filepath.Join(userFolder, "*.plan"))
	if err != nil {
		return err
	}
	if len(files) == 0 {
		return nil
	}
	return r.run("grep", append([]string{re}, files...)...)
}
