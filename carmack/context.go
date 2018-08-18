package carmack

import (
	"os/user"
	"path/filepath"

	"github.com/ruggi/carmack/git"
)

// Context contains runtime info and is passed to every command handler.
type Context struct {
	// Username is the current user's username, matching the plan files folder name.
	Username string
	// Folder is the main folder complete path (default: ~/.carmack).
	Folder string
}

// UserFolder returns the joined path of Folder and Username (e.g.: ~/.carmack/johndoe).
func (c *Context) UserFolder() string {
	return filepath.Join(c.Folder, c.Username)
}

// Load creates a new context.
func LoadContext(folderName string) (*Context, error) {
	u, err := user.Current()
	if err != nil {
		return nil, err
	}

	folder := filepath.Join(u.HomeDir, folderName)
	username := git.UserName(folder)
	if username == "" {
		username = u.Username
	}

	return &Context{
		Username: username,
		Folder:   folder,
	}, nil
}
