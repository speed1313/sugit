package main

import (
	"fmt"
	"github.com/speed1313/sugit/cmd"
	"os"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
        Name:  "sugit",
        Usage: "toy git",
		Commands: []*cli.Command{
			{
				Name: "init",
				Usage: "init git",
				Action: func(c *cli.Context) error {
					cmd.Git_init(os.Args[1:])
					return nil
				},
			},
			{
				Name: "add",
				Usage: "add file",
				Action: func(c *cli.Context) error {
					cmd.Git_add(os.Args[3:])
					return nil
				},
			},
			{
				Name: "cat-file",
				Usage: "cat file",
				Action: func(c *cli.Context) error {
					cmd.Git_cat_file(os.Args[3:])
					return nil
				},
			},
			{
				Name: "commit",
				Usage: "commit",
				Action: func(c *cli.Context) error {
					cmd.Git_commit(os.Args[3:])
					return nil
				},
			},
			{
				Name: "log",
				Usage: "log",
				Action: func(c *cli.Context) error {
					cmd.Git_log()
					return nil
				},
			},
		},
    }
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
	}
}
