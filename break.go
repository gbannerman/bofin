package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/nlopes/slack"
	"github.com/urfave/cli/v2"
)

func breakHandler(c *cli.Context, api *slack.Client, config *Config) {
	breakTime := 10

	if c.Args().Get(0) != "" {
		breakTime, _ = strconv.Atoi(c.Args().Get(0))
	}

	expiry := time.Now().Add(time.Duration(breakTime) * time.Minute)

	err := api.SetUserCustomStatusWithUser(config.UserID, "Taking a quick break", ":soon:", expiry.Unix())
	if err != nil {
		fmt.Printf("Error while updating status: %s\n", err)
		return
	}

	_, err = api.SetSnooze(breakTime)
	if err != nil {
		fmt.Printf("Error while setting snooze: %s\n", err)
		return
	}

	_, _, err = api.PostMessage(config.TeamChannelID, slack.MsgOptionText(fmt.Sprintf("Taking a %d minute break", breakTime), false), slack.MsgOptionAsUser(true))
	if err != nil {
		fmt.Printf("Error while posting message: %s\n", err)
		return
	}

	fmt.Printf("Starting a break for %d minutes\n", breakTime)
}

var Break = BofinCommand{
	Name:        "break",
	Usage:       "Updates Slack status to say you are taking a break",
	Description: "This command updates your Slack status to say you are taking a break, snoozes notifications and posts in your team channel",
	Handler:     breakHandler,
}
