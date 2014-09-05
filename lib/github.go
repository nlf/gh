package github

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type TokenRequest struct {
	Scopes []string `json:"scopes"`
	Note   string   `json:"note"`
}

type TokenResponse struct {
	Id    uint32 `json:"id"`
	Url   string `json:"url"`
	Token string `json:"token"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func (err *ErrorResponse) Error() string {
	return err.Message
}

func CreateToken(user string, password string, token string) (*TokenResponse, error) {
	client := &http.Client{}

	tokenRequest := &TokenRequest{Scopes: []string{"repo"}, Note: "go-gh extensions"}
	tokenJson, err := json.Marshal(tokenRequest)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "https://api.github.com/authorizations", bytes.NewReader(tokenJson))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(user, password)
	if token != "" {
		req.Header.Add("X-GitHub-OTP", token)
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		errorResp := &ErrorResponse{}
		json.NewDecoder(resp.Body).Decode(&errorResp)
		return nil, errorResp
	} else {
		tokenResp := &TokenResponse{}
		json.NewDecoder(resp.Body).Decode(&tokenResp)
		return tokenResp, nil
	}
}
