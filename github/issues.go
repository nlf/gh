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

type IssueParams struct {
	Milestone string
	State     string
	Assignee  string
	Creator   string
	Mentioned string
	Labels    []string
	Sort      string
	Direction string
	Since     string
}

type IssueSlice []Issue

func (client Client) GetIssues(repo string, params IssueParams) (IssueSlice, error) {
	query := make(map[string]string)

	query["state"] = params.State
	query["sort"] = params.Sort
	query["direction"] = params.Direction

	if len(params.Labels) > 0 {
		query["labels"] = strings.Join(params.Labels, ",")
	}

	if params.Milestone != "" {
		if params.Milestone == "none" || params.Milestone == "*" {
			query["milestone"] = params.Milestone
		} else {
			milestones, err := client.GetMilestones(repo, "all")
			if err != nil {
				log.Fatal(err)
			}

			for _, milestone := range milestones {
				if milestone.Title == params.Milestone {
					query["milestone"] = fmt.Sprintf("%d", milestone.Number)
					break
				}
			}
		}
	}

	if params.Assignee != "" {
		query["assignee"] = params.Assignee
	}

	if params.Creator != "" {
		query["creator"] = params.Creator
	}

	if params.Mentioned != "" {
		query["mentioned"] = params.Mentioned
	}

	if params.Since != "" {
		query["since"] = params.Since
	}

	issues := IssueSlice{}
	err := client.Request("GET", "/repos/"+repo+"/issues", query, nil, &issues)
	if err != nil {
		return nil, err
	}

	return issues, nil
}
