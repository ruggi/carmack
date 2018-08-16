package config

import "path/filepath"

// Config holds the main carmack configuration.
type Config struct {
	// Username is the current user's username, matching the plan files folder name.
	Username string `json:"username" yaml:"username" toml:"username"`
	// Folder is the main folder complete path (default: ~/.carmack).
	Folder string `json:"folder" yaml:"folder" toml:"folder"`
}

// UserFolder returns the joined path of Folder and Username (e.g.: ~/.carmack/johndoe).
func (c Config) UserFolder() string {
	return filepath.Join(c.Folder, c.Username)
}
