package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	Debug bool
	Core struct {
		DateInputLayout string `json:"DateInputLayout"`
	} `json:"Core"`
	Api struct {
		Nasdaq struct {
			ApiKey       string `json:"APIKey"`
			QueryColumns string `json:"QueryColumns"`
			DateFormat   string `json:"DateFormat"`
			EndpointUrl  string `json:"EndpointURL"`
		} `json:"Nasdaq"`
		Telegram struct {
			ApiKey          string `json:"APIKey"`
			ChannelId       string `json:"ChannelId"`
			ChannelUsername string `json:"ChannelUsername"`
			EndpointURL     string `json:"EndpointURL"`
			SendMessageSlug string `json:"SendMessageSlug"`
		} `json:"Telegram"`
	} `json:"Api"`
}

func GetConfig() Config {
	fileContent, err := os.ReadFile("./config.json")
	if err != nil {
		panic(err)
	}

	var config Config
	err = json.Unmarshal(fileContent, &config)
	if err != nil {
		panic(err)
	}

	return config
}
