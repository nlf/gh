package github

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func (client *Client) GetIssues(repo string) ([]Issue, error) {
	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", client.BaseURL+"/repos/"+repo+"/issues", nil)
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
		json.NewDecoder(resp.Body).Decode(&errorResp)
		return nil, errorResp
	}

	issues := &[]Issue{}
	json.NewDecoder(resp.Body).Decode(&issues)
	return *issues, nil
}

// This method doesn't use the normal client since it is used to generate the config
func CreateToken(baseUrl string, user string, password string, otp string) (TokenResponse, error) {
	client := &http.Client{}

	tokenRequest := &TokenRequest{Scopes: []string{"repo"}, Note: "go-gh extensions"}
	tokenJson, err := json.Marshal(tokenRequest)
	if err != nil {
		return *&TokenResponse{}, err
	}

	req, err := http.NewRequest("POST", baseUrl+"/authorizations", bytes.NewReader(tokenJson))
	if err != nil {
		return *&TokenResponse{}, err
	}

	req.SetBasicAuth(user, password)
	if otp != "" {
		req.Header.Add("X-GitHub-OTP", otp)
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return *&TokenResponse{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		errorResp := &ErrorResponse{}
		json.NewDecoder(resp.Body).Decode(&errorResp)
		return *&TokenResponse{}, errorResp
	}

	tokenResp := &TokenResponse{}
	json.NewDecoder(resp.Body).Decode(&tokenResp)
	return *tokenResp, nil
}
