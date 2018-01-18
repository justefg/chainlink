package main

import (
	"os"

	"github.com/smartcontractkit/chainlink/cmd"
	"github.com/smartcontractkit/chainlink/store"
	"github.com/urfave/cli"
)

func main() {
	client := cmd.Client{cmd.RendererTable{os.Stdout}, store.NewConfig()}
	app := cli.NewApp()
	app.Usage = "CLI for Chainlink"
	app.Commands = []cli.Command{
		{
			Name:    "node",
			Aliases: []string{"n"},
			Usage:   "Run the chainlink node",
			Action:  client.RunNode,
		},
		{
			Name:    "jobs",
			Aliases: []string{"j"},
			Usage:   "Get all jobs",
			Action:  client.GetJobs,
		},
		{
			Name:    "show",
			Aliases: []string{"s"},
			Usage:   "Show a specific job",
			Action:  client.ShowJob,
		},
	}
	app.Run(os.Args)
}
