package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/deadcheat/gengou"
	"github.com/gorilla/mux"
	"gopkg.in/urfave/cli.v2"
)

func main() {
	app := cli.App{
		Name:    "gengou",
		Version: "0.1.0",
		Commands: []*cli.Command{
			&cli.Command{
				Name:    "standalone",
				Aliases: []string{"a"},
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
			&cli.Command{
				Name:    "server",
				Aliases: []string{"s"},
				Usage:   "start gengou server that serves JSON Response with /year/month/day url",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:        "host",
						DefaultText: "localhost",
						Usage:       "server host",
					},
					&cli.IntFlag{
						Name:        "port",
						DefaultText: "8080",
						Usage:       "server port",
					},
					&cli.StringFlag{
						Name:        "baseuri",
						DefaultText: "/",
						Usage:       "base uri",
					},
				},
				Action: func(c *cli.Context) error {
					r := mux.NewRouter()
					r.HandleFunc("/{year:[0-9]+}/{month:[0-9]+}/{day:[0-9]+}", func(http.ResponseWriter, *http.Request) {

					})
					http.Handle("/", r)
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
