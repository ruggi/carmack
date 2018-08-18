# carmack

[![Go Report Card](https://goreportcard.com/badge/github.com/ruggi/carmack)](https://goreportcard.com/report/github.com/ruggi/carmack)

carmack is a daily progress tracker for teams, inspired by [John Carmack's .plan files](https://github.com/ESWAT/john-carmack-plan-archive) and based on git.

Daily notes and tasks are stored in plain `.plan` files, using a simple and grep-able [format](#plan-format), divided into *done*, *completed*, *canceled*, and *open* entries.

carmack stores each user's (or teammate!) daily plan files in [subfolders](#structure), so it's easy to see who did what on a given date.

The data is stored in a git repository keeping track of everyone's activity.

## Installation

```sh
go get github.com/ruggi/carmack
```

## Usage

### Adding entries

You can use the `add` command to add new entries.

Mark them as *done*, *completed*, or *canceled* tasks with the `--done`, `--completed`, or `--canceled` flags.

```sh
$ carmack add 'create PR for new feature'
```

```sh
$ carmack add --done 'merged feature branch'
$ carmack add --completed 'cleanup code'
$ carmack add --canceled 'weird feature'
```

### Showing entries

You can use the `show` command to show entries from plan files. They can be filtered with the `--done`, `--completed`, `--canceled`, or `--open` flags.

By default it shows entries for the your (current user's) plan files.

It's possible to use the `--user [username]` option to show entries for a specific user.

```sh
$ carmack show [-u username] [--done|--completed|--canceled|--open]
```

The nice thing about the plan files is that they're just plain text, so they can be directly grep-ed, modified, deleted, and used with other cli tools.

### Listing plan files

You can use the `list` (alias: `ls`) command to get a list of your plan files.

```sh
$ carmack list
```

### Git

The main benefit of using carmack for teams is that synchronization is done using git, with the `~/.carmack` folder being a repository itself.

First, initialize the git repo:

```sh
$ carmack git init

Initialized empty Git repository in /Users/john/.carmack/.git/
[master (root-commit) c6f290d] initial commit
 1 file changed, 17 insertions(+)
 create mode 100755 john/2018-08-17.plan
```

Then, set up the remotes with your team's carmack repository:

```sh
$ carmack git remote add origin git@github.com:team/plan
```

Now you can sync your carmack folder with `pull` and `push`:

```sh
$ carmack git pull origin master
$ carmack git push -u origin master
```

Once the git repository is initialized, every time you add a new entry to your plan files it gets
automatically added and committed to the repo:

```sh
$ carmack add -d 'closed issue'
[master 9a13e06] 'john: plan update 2018-08-17T22:02:40+02:00'
 1 file changed, 1 insertion(+)

$ carmack git pull
$ carmack git push
```

You can explicitly set a local git `user.name` config variable on the git repository so that your entries will be stored
in a folder with that name:

```sh
$ carmack git config --local user.name '<your_name>'
```

#### Other git commands

`carmack git` is just a wrapper around git itself, so you can use any normal command with it (a-là [pass](https://www.passwordstore.org/)).

For example:
```sh
$ carmack git log --oneline
6fa0ca9 (HEAD -> master) 'john: plan update 2018-08-17T22:02:40+02:00'
9a13e06 'bob: plan update 2018-08-17T21:05:37+02:00'
5b3de41 initial commit
```

## Structure
All plan files are stored in the `~/.carmack/` directory (and git repo), and each user's own plan files are stored in the `~/.carmack/<user>/` folders.

```sh
~/.carmack
├── alice
│   ├── 2018-08-15.plan
│   └── 2018-08-17.plan
├── bob
│   ├── 2018-08-14.plan
│   ├── 2018-08-15.plan
│   ├── 2018-08-16.plan
│   └── 2018-08-17.plan
└── john
    ├── 2018-08-15.plan
    ├── 2018-08-16.plan
    └── 2018-08-17.plan
```

When a new entry is added, the corresponding plan file is created, if it doesn't exist, with the name `YYYY-MM-DD.plan`. The file name's date format is UTC-based.

### User folder names

The folder where a user's plan files are stored is named as:

* the local (to the carmack repo) git `user.name` config value, or
* the global git `user.name` config value, or
* the current login username (as in `$USER` or `whoami`)

## Plan format

As nicely written [here](https://garbagecollected.org/2017/10/24/the-carmack-plan/), a plan file keeps track of daily work.

```plan
* merged PR
* answered that_guy on intercom
* deployed version 42

+ reviewed bob's feature branch

- meeting with that_guy

fix cloudformation
add more workers to the async stuff
```

The syntax is very simple:

| Prefix    | Meaning                                         |
|-----------| ------------------------------------------------|
| `*`       | Completed on the same day                       |
| `+`       | Completed on a later day                        |
| `-`       | Canceled/decided against on a later day         |
| No prefix | Open/Note                                       |

When parsing and saving a plan file, carmack will adjust its content nicely, so that entries with the same prefix are grouped together.
