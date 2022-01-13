package main

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type WeblinkConfig struct {
	WeblinkPortal string `yaml:"weblink_portal"`
	RefreshToken  string `yaml:"refresh_token"`
	LongLifeToken string `yaml:"long_life_token"`
}

type GeneralConfig struct {
	ClientID      string `yaml:"adm_client_id"`
	ClientSecret  string `yaml:"adm_client_secret"`
	BootDirectory string `yaml:"boot_dir"`
}

type BootEnvConfig struct {
	General GeneralConfig            `yaml:"general"`
	Weblink map[string]WeblinkConfig `yaml:"weblink"`
}

type Config struct {
	APIToken      string        `yaml:"slack_api_token"`
	UserID        string        `yaml:"slack_user_id"`
	TeamChannelID string        `yaml:"slack_team_channel_id"`
	BootEnvConfig BootEnvConfig `yaml:"boot_env"`
}

func GetConfig() *Config {
	config := Config{}

	configPath, err := os.UserHomeDir()
	if err != nil {
		log.Printf("Home directory error: #%v ", err)
	}

	configPath = configPath + "/.bofin-conf.yaml"

	yamlFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return &config
}
