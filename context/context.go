package context

import "path/filepath"

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
