package main

import (
	"fmt"
	"os/exec"
	"reflect"

	"github.com/nlopes/slack"
	"github.com/urfave/cli/v2"
)

func weblinkHandler(c *cli.Context, api *slack.Client, config *Config) {
	type Env struct {
		ApiAccessToken       string `bootvar:"API_ACCESS_TOKEN"`
		WeblinkStagingDomain string `bootvar:"WEBLINK_STAGING_DOMAIN"`
	}

	instance := "lmsdemo"

	if c.Args().Get(0) != "" {
		instance = c.Args().Get(0)
	}

	env := Env{
		config.BootEnvConfig.Weblink[instance].LongLifeToken,
		config.BootEnvConfig.Weblink[instance].WeblinkPortal,
	}

	if env.WeblinkStagingDomain == "" {
		fmt.Printf("Error: No weblink config found for instance `%s`\n", instance)
		return
	}

	if env.ApiAccessToken == "" {
		resp, _ := RefreshCoreToken(config, instance)

		env.ApiAccessToken = resp.AccessToken
	}

	ReplaceEnv(config, reflect.TypeOf(env), reflect.ValueOf(env))

	cmd := exec.Command("make", "weblink")
	cmd.Dir = config.BootEnvConfig.General.BootDirectory
	cmd.Start()

	fmt.Printf("Updated boot weblink config to point at `%s`\n", instance)
}

var Weblink = BofinCommand{
	Name:        "weblink",
	Usage:       "Updates boot .env file to point at a given weblink portal",
	Description: "This command updates your boot .env file to point at a given weblink portal and restarts the relevant containers",
	Handler:     weblinkHandler,
}
