package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/mickael/zuo/internal/command"
	"github.com/mickael/zuo/internal/print"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "zuo",
		Usage: "Command-line interface to Zuora",
		Commands: []*cli.Command{
			{
				Name:    "query",
				Aliases: []string{"q"},
				Usage:   "Queries a Zuora object by its ID",
				Action: func(c *cli.Context) error {
					fmt.Printf(
						"querying %s with id %s\n",
						c.Args().Get(0),
						c.Args().Get(1),
					)
					return nil
				},
			},
			{
				Name:  "exec",
				Usage: "Executes a ZOQL query",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:  "json",
						Usage: "Raw JSON from Zuora",
					},
				},
				Action: func(c *cli.Context) error {
					queryString := c.Args().Get(0)
					if len(queryString) == 0 {
						fmt.Println("query is required")
						return nil
					}

					resp := command.Query(queryString)

					if c.Bool("json") {
						json, err := json.Marshal(resp)
						if err != nil {
							log.Fatal(err)
						}
						fmt.Println(string(json))
					} else {
						fmt.Printf(
							"Found %d record(s)\n",
							resp.Size,
						)
						print.Table(resp.Records)
					}
					return nil
				},
			},
			{
				Name:  "login",
				Usage: "Generates an access token from Zuora",
				Action: func(c *cli.Context) error {
					fmt.Println("login...")
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
