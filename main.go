package main

import (
	"os"
	"strings"

	"github.com/ruggi/carmack/carmack"
	"github.com/ruggi/carmack/commands"
	"github.com/ruggi/carmack/plan"
	"github.com/urfave/cli"
)

const (
	folderName = ".carmack"
)

func main() {
	ctx, err := carmack.LoadContext(folderName)
	if err != nil {
		panic(err)
	}
	err = os.MkdirAll(ctx.UserFolder(), 0755)
	if err != nil {
		panic(err)
	}

	app := cli.NewApp()
	app.Name = "carmack"
	app.Usage = "track daily progress with .plan files"
	app.Version = carmack.Version
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
					Usage: "[-] for something old that has been canceled/decided against",
				},
			},
			Action: func(c *cli.Context) error {
				entry := strings.Join(c.Args(), " ")
				entryType := makeEntryTypeFromFlags(c)
				return commands.Add(ctx, entry, entryType)
			},
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
			Action: func(c *cli.Context) error {
				user := c.String("user")
				entryType := makeEntryTypeFromFlags(c)
				return commands.Show(ctx, user, entryType)
			},
		},
		{
			Name:    "list",
			Aliases: []string{"ls"},
			Usage:   "List all plan files",
			Action:  commands.List(ctx),
		},
		{
			Name:  "git",
			Usage: "Issue git commands",
			Subcommands: []cli.Command{
				{
					Name:  "init",
					Usage: "initialize plan files git repo",
					Action: func(c *cli.Context) error {
						return commands.GitInit(ctx)
					},
				},
			},
			SkipFlagParsing: true,
			Action: func(c *cli.Context) error {
				return commands.Git(ctx, c.Args()...)
			},
		},
	}

	app.RunAndExitOnError()
}

func makeEntryTypeFromFlags(c *cli.Context) plan.EntryType {
	if c.Bool("done") {
		return plan.Done
	}
	if c.Bool("completed") {
		return plan.Completed
	}
	if c.Bool("canceled") {
		return plan.Canceled
	}
	return plan.Note
}
