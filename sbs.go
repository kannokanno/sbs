package main

import (
	"github.com/codegangsta/cli"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "boom"
	app.Usage = "sample"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "lang, l",
			Value: "english",
			Usage: "hogehoge",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:    "add",
			Aliases: []string{"a"},
			Usage:   "add a task",
			Action: func(c *cli.Context) {
				if len(c.Args()) > 0 {
					println("yeah")
				} else {
					println(c.String("lang"))
				}
			},
		},
	}
	app.Run(os.Args)
}
