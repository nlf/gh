package github

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type TokenRequest struct {
	Scopes       []string `json:"scopes,omitempty"`
	Note         string   `json:"note,omitempty"`
	NoteURL      string   `json:"note_url,omitempty"`
	ClientID     string   `json:"client_id,omitempty"`
	ClientSecret string   `json:"client_secret,omitempty"`
}

type Token struct {
	Id     uint32   `json:"id"`
	URL    string   `json:"url"`
	Scopes []string `json:"scopes"`
	Token  string   `json:"token"`
	App    struct {
		URL      string `json:"url"`
		Name     string `json:"name"`
		ClientID string `json:"name"`
	} `json:"app"`
	Note    string `json:"note"`
	NoteURL string `json:"note_url"`
	Updated string `json:"updated_at"`
	Created string `json:"created_at"`
}

func (client Client) CreateToken(user string, password string, otp string) (Token, error) {
	httpClient := http.Client{}

	tokenRequest := TokenRequest{Scopes: []string{"repo"}, Note: "go-gh extensions"}
	tokenJson, err := json.Marshal(tokenRequest)
	if err != nil {
		return Token{}, err
	}

	req, err := http.NewRequest("POST", client.BaseURL+"/authorizations", bytes.NewReader(tokenJson))
	if err != nil {
		return Token{}, err
	}

	req.SetBasicAuth(user, password)
	if otp != "" {
		req.Header.Add("X-GitHub-OTP", otp)
	}

	req.Header.Add("Content-Type", "application/json")

	resp, err := httpClient.Do(req)
	if err != nil {
		return Token{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != 201 {
		errorResp := &ErrorResponse{}
		json.NewDecoder(resp.Body).Decode(errorResp)
		return Token{}, errorResp
	}

	tokenResp := Token{}
	json.NewDecoder(resp.Body).Decode(&tokenResp)
	return tokenResp, nil
}
