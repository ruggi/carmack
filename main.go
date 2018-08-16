package main

import (
	"os"
	"os/user"
	"path/filepath"

	"github.com/ruggi/carmack/commands"
	"github.com/ruggi/carmack/config"
	"github.com/urfave/cli"
)

const (
	folderName = ".carmack"
)

var (
	// Version is the release version semver number.
	Version = "1.0.0"
)

func main() {
	cfg, err := setupConfig()
	if err != nil {
		panic(err)
	}
	err = os.MkdirAll(cfg.UserFolder(), 0755)
	if err != nil {
		panic(err)
	}

	app := cli.NewApp()
	app.Name = "carmack"
	app.Usage = "track daily progress with .plan files"
	app.Version = Version
	app.Author = "Federico Ruggi"
	app.Description = "// TODO //"
	app.Commands = []cli.Command{
		{
			Name:  "add",
			Usage: "Add a new item to today's plan",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "done,d",
					Usage: "[*] for something you started and completed today",
				},
				cli.BoolFlag{
					Name:  "completed,c",
					Usage: "[+] for something old that you completed today",
				},
				cli.BoolFlag{
					Name:  "canceled,x",
					Usage: "[-] for something old that has been canceled",
				},
			},
			Action: commands.Add(cfg.Username, cfg.Folder, cfg.UserFolder()),
		},
		{
			Name:   "open",
			Usage:  "Show open items",
			Action: commands.Open(cfg.UserFolder()),
		},
		{
			Name:    "list",
			Aliases: []string{"ls"},
			Usage:   "List all plan files",
			Action:  commands.List(cfg.UserFolder()),
		},
		{
			Name:  "git",
			Usage: "Issue git commands",
			Subcommands: []cli.Command{
				{
					Name:   "init",
					Usage:  "initialize plan files git repo",
					Action: commands.GitInit(cfg.Folder),
				},
			},
			SkipFlagParsing: true,
			Action:          commands.Git(cfg.Folder),
		},
	}
	app.RunAndExitOnError()
}

func setupConfig() (config.Config, error) {
	u, err := user.Current()
	if err != nil {
		return config.Config{}, err
	}
	return config.Config{
		Username: u.Username,
		Folder:   filepath.Join(u.HomeDir, folderName),
	}, nil
}
