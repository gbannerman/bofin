package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/nlopes/slack"
	"github.com/urfave/cli/v2"
)

func focusHandler(c *cli.Context, api *slack.Client, config *Config) {
	focusTime := 30

	if c.Args().Get(0) != "" {
		focusTime, _ = strconv.Atoi(c.Args().Get(0))
	}

	expiry := time.Now().Add(time.Duration(focusTime) * time.Minute)

	err := api.SetUserCustomStatusWithUser(config.UserID, "Focusing - available via email or zoom", ":male-technologist:", expiry.Unix())
	if err != nil {
		fmt.Printf("Error while updating status: %s\n", err)
		return
	}

	_, err = api.SetSnooze(focusTime)
	if err != nil {
		fmt.Printf("Error while setting snooze: %s\n", err)
		return
	}

	fmt.Printf("Focus enabled for %d minutes\n", focusTime)
}

var Focus = BofinCommand{
	Name:        "focus",
	Usage:       "Updates Slack status to say you are focusing",
	Description: "This command updates your Slack status to say you are focusing and also snoozes notifications",
	Handler: focusHandler,
}