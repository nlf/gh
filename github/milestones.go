package github

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

type Milestone struct {
	URL          string `json:"url"`
	Number       uint32 `json:"number"`
	State        string `json:"state"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	Creator      User   `json:"creator"`
	OpenIssues   uint32 `json:"open_issues"`
	ClosedIssues uint32 `json:"closed_issues"`
	Created      string `json:"created_at"`
	Updated      string `json:"updated_at"`
	DueOn        string `json:"due_on"`
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
