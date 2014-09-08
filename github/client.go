package github

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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

func (client Client) Request(method string, path string, query map[string]string, payload []byte, receiver interface{}) error {
	httpClient := http.Client{}
	u, err := url.Parse(client.BaseURL)
	if err != nil {
		log.Fatal(err)
	}

	u.Path += path
	q := u.Query()

	for key, val := range query {
		q.Add(key, val)
	}

	u.RawQuery = q.Encode()

	req, err := http.NewRequest(method, u.String(), bytes.NewReader(payload))
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "gh")
	req.Header.Add("Authorization", "token "+client.Token)

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		errorResp := &ErrorResponse{}
		json.NewDecoder(resp.Body).Decode(errorResp)
		return errorResp
	}

	json.NewDecoder(resp.Body).Decode(receiver)
	return nil
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
