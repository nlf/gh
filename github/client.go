package github

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os/user"
)

type Client struct {
	BaseURL string `json:"baseUrl"`
	Token   string `json:"token"`
}

type ErrorResponse struct {
	Message          string `json:"message"`
	DocumentationURL string `json:"documentation_url"`
}

func (err *ErrorResponse) Error() string {
	return err.Message
}

func GetClient() Client {
	client := Client{}
	client.LoadConfig()

	return client
}

func (client *Client) LoadConfig() {
	configPath := GetConfigPath()
	configJson, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(configJson, &client)
	if err != nil {
		log.Fatal(err)
	}
}

func (client Client) SaveConfig() {
	configPath := GetConfigPath()
	configJson, err := json.Marshal(client)
	if err != nil {
		log.Fatal(err)
	}

	err = ioutil.WriteFile(configPath, configJson, 0600)
	if err != nil {
		log.Fatal(err)
	}
}

func GetConfigPath() string {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	configPath := currentUser.HomeDir + "/.gh"
	return configPath
}
