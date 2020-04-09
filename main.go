package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mickael/zuo/internal/command"
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
				Action: func(c *cli.Context) error {
					queryString := c.Args().Get(0)
					response := command.Query(queryString)
					fmt.Printf("%+v\n", response)
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
