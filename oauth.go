package main

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type RefreshResponse struct {
	AccessToken string `json:"access_token"`
}

func RefreshCoreToken(config *Config, instance string) (*RefreshResponse, error) {
	var refreshToken = config.BootEnvConfig.Weblink[instance].RefreshToken

	data := url.Values{
		"grant_type":    {"refresh_token"},
		"client_id":     {config.BootEnvConfig.General.ClientID},
		"client_secret": {config.BootEnvConfig.General.ClientSecret},
		"refresh_token": {refreshToken},
	}

	resp, _ := http.PostForm("https://auth.stagingadministratehq.com/oauth/token", data)

	var response RefreshResponse

	err := json.NewDecoder(resp.Body).Decode(&response)
	return &response, err
}
