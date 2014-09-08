package github

import (
	"fmt"
	"log"
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
	query := make(map[string]string)
	if len(labels) > 0 {
		query["labels"] = strings.Join(labels, ",")
	}

	if milestone != "" {
		if milestone == "none" || milestone == "*" {
			query["milestone"] = milestone
		} else {
			milestones, err := client.GetMilestones(repo, "all")
			if err != nil {
				log.Fatal(err)
			}

			for _, mstone := range milestones {
				if mstone.Title == milestone {
					query["milestone"] = fmt.Sprintf("%d", mstone.Number)
					break
				}
			}
		}
	}

	issues := IssueSlice{}
	err := client.Request("GET", "/repos/"+repo+"/issues", query, nil, &issues)
	if err != nil {
		return nil, err
	}

	return issues, nil
}
