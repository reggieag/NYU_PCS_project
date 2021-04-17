package main

import (
	"io/ioutil"

	"github.com/go-oauth2/oauth2/v4/models"
	"github.com/go-oauth2/oauth2/v4/store"
	"gopkg.in/yaml.v3"
)

type Client struct {
	ClientID     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
}

type ClientConfig struct {
	Clients []Client `yaml:"clients"`
}

func readConfig(filename string) (ClientConfig, error) {
	var config ClientConfig
	configBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return config, err
	}
	if err := yaml.Unmarshal(configBytes, &config); err != nil {
		return config, err
	}
	return config, nil
}

func createClients(clientStore *store.ClientStore, config ClientConfig) {
	for _, client := range config.Clients {
		clientStore.Set(client.ClientID, &models.Client{
			ID:     client.ClientID,
			Secret: client.ClientSecret,
			Domain: "http://127.0.0.1",
		})
	}
}
