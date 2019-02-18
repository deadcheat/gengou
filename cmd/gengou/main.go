package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/deadcheat/gengou"
	"gopkg.in/urfave/cli.v2"
)

func main() {
	app := cli.App{
		Name:    "gengou",
		Version: "0.0.1",
		Commands: []*cli.Command{
			&cli.Command{
				Name:    "standalone",
				Aliases: []string{"s"},
				Usage:   "give me year/month/day formatted string, gengou answers gengou to you!!",
				Action: func(c *cli.Context) error {
					args := c.Args().Slice()
					if len(args) < 1 {
						return errors.New("no args are found")
					}

					gengos := make([]gengou.Gengou, 0)
					for i := range args {
						pgs, err := gengou.Find(args[i])
						if err != nil {
							fmt.Println(err)
							continue
						}
						gengos = append(gengos, pgs...)
					}

					for i := range gengos {
						fmt.Println(gengos[i].Name)
					}

					return nil
				},
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}
