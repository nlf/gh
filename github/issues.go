package github

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Issue struct {
	URL         string      `json:"url"`
	HtmlURL     string      `json:"html_url"`
	Number      uint32      `json:"number"`
	State       string      `json:"state"`
	Title       string      `json:"title"`
	Body        string      `json:"body"`
	User        User        `json:"user"`
	Labels      []Label     `json:"labels"`
	Assignee    User        `json:"assignee"`
	Milestone   Milestone   `json:"milestone"`
	Comments    uint32      `json:"comments"`
	PullRequest PullRequest `json:"pull_request"`
	Closed      string      `json:"closed_at"`
	Created     string      `json:"created_at"`
	Updated     string      `json:"updated_at"`
}

type IssueSlice []Issue

func (client Client) GetIssues(repo string, labels []string, milestone string) (IssueSlice, error) {
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

	issues := IssueSlice{}
	json.NewDecoder(resp.Body).Decode(&issues)
	return issues, nil
}
