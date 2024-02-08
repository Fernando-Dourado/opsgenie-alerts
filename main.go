package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Drone-OpsGenie-Alert"
	app.Action = run
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:     "token",
			Usage:    "OpsGenie API key",
			EnvVar:   "PLUGIN_OPSGENIE_TOKEN",
			Required: true,
		},
		cli.StringFlag{
			Name:     "message",
			Usage:    "Alert message content",
			EnvVar:   "PLUGIN_OPSGENIE_MESSAGE",
			Required: true,
		},
		cli.StringFlag{
			Name:     "description",
			Usage:    "Alert description content",
			EnvVar:   "PLUGIN_OPSGENIE_DESCRIPTION",
			Required: false,
		},
	}
	app.Run(os.Args)
}

func run(c *cli.Context) {
	plugin := Plugin{
		Config: Config{
			Token: c.String("token"),
		},
		Alert: Alert{
			Message:     c.String("message"),
			Description: c.String("description"),
		},
	}

	if err := plugin.Exec(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Result <%s>", plugin.Result)
}
