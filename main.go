package main

import (
	"os"
	"os/user"
	"path/filepath"

	"github.com/ruggi/carmack/commands"
	"github.com/ruggi/carmack/config"
	"github.com/ruggi/carmack/git"
	"github.com/ruggi/carmack/util"
	"github.com/urfave/cli"
)

const (
	folderName = ".carmack"
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
	app.Version = util.Version
	app.Author = "Federico Ruggi"
	app.Description = "// TODO //"
	app.Commands = []cli.Command{
		{
			Name:  "add",
			Usage: "Add a new entry to today's plan",
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
			Name:  "show",
			Usage: "Show open entries",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "open,o",
					Usage: "show open entries",
				},
				cli.BoolFlag{
					Name:  "done,d",
					Usage: "show done entries",
				},
				cli.BoolFlag{
					Name:  "completed,c",
					Usage: "show completed entries",
				},
				cli.BoolFlag{
					Name:  "canceled,x",
					Usage: "show canceled entries",
				},
				cli.StringFlag{
					Name:  "user,u",
					Usage: "show entries for a specific user",
				},
			},
			Action: commands.Show(cfg.Folder, cfg.UserFolder()),
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
	folder := filepath.Join(u.HomeDir, folderName)
	username := git.UserName(folder)
	if username == "" {
		username = u.Username
	}
	return config.Config{
		Username: username,
		Folder:   folder,
	}, nil
}
