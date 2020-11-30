package main

import (
	"log"
	"os"
	"sort"
	"fmt"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "gbuilder"
	app.Usage = "go project structure build tool"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:     "name, n",
			Usage:    "project name",
			Required: true,
		},

		cli.StringFlag{
			Name:     "author, a",
			Usage:    "project author",
			Required: true,
		},
	}

	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))
	app.Action = func(ctx *cli.Context) error {

		name := ctx.String("name")
		author := ctx.String("author")

		builder := NewBuilder(name, author, os.ModePerm)
		if err := builder.Run(); err != nil {
			os.Exit(1)
		}

		return nil
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		log.Fatal(err.Error())
	}
}
