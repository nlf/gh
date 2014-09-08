package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func GetClient() Client {
	client := Client{}
	client.LoadConfig()

	return client
}

func (client Client) GetMilestones(repo string, state string) ([]Milestone, error) {
	httpClient := http.Client{}
	u, err := url.Parse(client.BaseURL)
	if err != nil {
		log.Fatal(err)
	}

	u.Path = u.Path + "/repos/" + repo + "/milestones"
	if state != "" {
		u.Query().Add("state", state)
	}

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "go-gh")
	req.Header.Add("Authorization", "token "+client.Token)

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		errorResp := &ErrorResponse{}
		json.NewDecoder(resp.Body).Decode(errorResp)
		return nil, errorResp
	}

	milestones := []Milestone{}
	json.NewDecoder(resp.Body).Decode(&milestones)
	return milestones, nil
}

func (client Client) GetIssues(repo string, labels []string, milestone string) ([]Issue, error) {
	httpClient := http.Client{}
	u, err := url.Parse(client.BaseURL)
	if err != nil {
		log.Fatal(err)
	}

	u.Path = u.Path + "/repos/" + repo + "/issues"
	query := u.Query()

	if len(labels) > 0 {
		query.Add("labels", strings.Join(labels, ","))
	}

	if milestone != "" {
		if milestone == "none" || milestone == "*" {
			query.Add("milestone", milestone)
		} else {
			milestones, err := client.GetMilestones(repo, "all")
			if err != nil {
				log.Fatal(err)
			}

			for _, mstone := range milestones {
				if mstone.Title == milestone {
					query.Add("milestone", fmt.Sprintf("%d", mstone.Number))
					break
				}
			}
		}
	}

	u.RawQuery = query.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("User-Agent", "go-gh")
	req.Header.Add("Authorization", "token "+client.Token)

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		errorResp := &ErrorResponse{}
		json.NewDecoder(resp.Body).Decode(errorResp)
		return nil, errorResp
	}

	issues := []Issue{}
	json.NewDecoder(resp.Body).Decode(&issues)
	return issues, nil
}

// This method doesn't use the normal client since it is used to generate the config
func CreateToken(baseUrl string, user string, password string, otp string) (TokenResponse, error) {
	client := http.Client{}

	tokenRequest := TokenRequest{Scopes: []string{"repo"}, Note: "go-gh extensions"}
	tokenJson, err := json.Marshal(tokenRequest)
	if err != nil {
		return TokenResponse{}, err
	}

	req, err := http.NewRequest("POST", baseUrl+"/authorizations", bytes.NewReader(tokenJson))
	if err != nil {
		return TokenResponse{}, err
	}

	req.SetBasicAuth(user, password)
	if otp != "" {
		req.Header.Add("X-GitHub-OTP", otp)
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return TokenResponse{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		errorResp := &ErrorResponse{}
		json.NewDecoder(resp.Body).Decode(errorResp)
		return TokenResponse{}, errorResp
	}

	tokenResp := TokenResponse{}
	json.NewDecoder(resp.Body).Decode(&tokenResp)
	return tokenResp, nil
}
