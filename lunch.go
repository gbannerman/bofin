package main

import (
	"fmt"
	"time"

	"github.com/nlopes/slack"
	"github.com/urfave/cli/v2"
)

func lunchHandler(c *cli.Context, api *slack.Client, config *Config) {
	lunchLength := 45

	expiry := time.Now().Add(time.Duration(lunchLength) * time.Minute)

	err := api.SetUserCustomStatusWithUser(config.UserID, "Lunch - will reply when I'm back", ":pizza:", expiry.Unix())
	if err != nil {
		fmt.Printf("Error while updating status: %s\n", err)
		return
	}

	_, err = api.SetSnooze(lunchLength)
	if err != nil {
		fmt.Printf("Error while setting snooze: %s\n", err)
		return
	}

	_, _, err = api.PostMessage(config.TeamChannelID, slack.MsgOptionText("Grabbing lunch :pizza:", false), slack.MsgOptionAsUser(true))
	if err != nil {
		fmt.Printf("Error while posting message: %s\n", err)
		return
	}

	fmt.Print("Lunch started\n")
}

var Lunch = BofinCommand{
	Name:        "lunch",
	Usage:       "Updates Slack status to say you are taking lunch",
	Description: "This command updates your Slack status to say you are at lunch, snoozes notifications and posts in your team channel",
	Handler: lunchHandler,
}