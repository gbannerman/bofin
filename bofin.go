package main

import (
	"os"

	"github.com/nlopes/slack"
	"github.com/urfave/cli/v2"
)

type Handler func(*cli.Context, *slack.Client, *Config)

type BofinCommand struct {
	Name string
	Usage string
	Description string
	Handler Handler
}

var bofinCommands = []BofinCommand{
	Weblink,
	Focus,
	Lunch,
	Break,
}

func main() {
	config := GetConfig()

	api := slack.New(config.APIToken)

	app := cli.NewApp()
	app.Name = "bofin"
	app.Usage = "A CLI tool for the Booking & Finance tribe to automate work tasks"
	app.UsageText = "bofin command [arguments...]"
	app.EnableBashCompletion = true

	var commands = []*cli.Command{}

	for i := range bofinCommands {
		i := i
		commands = append(commands, &cli.Command{
			Name: bofinCommands[i].Name,
			Usage: bofinCommands[i].Usage,
			Description: bofinCommands[i].Description,
			Action: func(c *cli.Context) error {
				handler := bofinCommands[i].Handler
				handler(c, api, config)
				return nil
			},
		})
    }

	app.Commands = commands

	_ = app.Run(os.Args)

}
