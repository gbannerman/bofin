package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"

	"text/template"

	"github.com/nlopes/slack"
	"github.com/urfave/cli/v2"
)

func weblinkHandler(c *cli.Context, api *slack.Client, config *Config) {
	type Env struct {
		ApiAccessToken       string
		WeblinkStagingDomain string
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
		fmt.Printf("Error: No bootenv config found for instance `%s`\n", instance)
		return
	}

	if env.ApiAccessToken == "" {
		resp, _ := refreshCoreToken(config, instance)

		env.ApiAccessToken = resp.AccessToken
	}

	envTemplatePath := fmt.Sprintf("%s/.env.template", config.BootEnvConfig.General.BootDirectory)
	envTemplate, _ := ioutil.ReadFile(envTemplatePath)

	tmpl, _ := template.New("env").Parse(string(envTemplate))

	envFilePath := fmt.Sprintf("%s/.env", config.BootEnvConfig.General.BootDirectory)
	envFile, _ := os.OpenFile(envFilePath, os.O_WRONLY, 0755)

	tmpl.Execute(envFile, env)

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
