package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/codegangsta/cli"
	"github.com/nlf/gh/git"
	"github.com/nlf/gh/github"
)

func PrintHeader() {
	fmt.Printf("\x1b[1m%-5s %-80.80s %-15.15s %-15.15s %-15.15s %-5s\x1b[0m\n", "ID", "Title", "Assignee", "Labels", "Milestone", "Comments")
}

func PrintIssue(issue github.Issue, bold bool) {
	pullRequest := issue.PullRequest.URL != ""
	lineFormat := ""

	if bold {
		if pullRequest {
			lineFormat = "\x1b[1m%-5d [pr] %-75.75s %-15.15s %-15.15s %-15.15s %-5d\x1b[0m\n"
		} else {
			lineFormat = "\x1b[1m%-5d %-80.80s %-15.15s %-15.15s %-15.15s %-5d\x1b[0m\n"
		}
	} else {
		if pullRequest {
			lineFormat = "%-5d [pr] %-75.75s %-15.15s %-15.15s %-15.15s %-5d\n"
		} else {
			lineFormat = "%-5d %-80.80s %-15.15s %-15.15s %-15.15s %-5d\n"
		}
	}

	labels := make([]string, 0)
	for _, label := range issue.Labels {
		labels = append(labels, label.Name)
	}

	fmt.Printf(lineFormat, issue.Number, issue.Title, issue.Assignee.Login, strings.Join(labels, ","), issue.Milestone.Title, issue.Comments)
}

func Issues(c *cli.Context) {
	client := github.GetClient()
	repo := git.GetRepo()

	params := github.IssueParams{
		Milestone: c.String("milestone"),
		State:     c.String("state"),
		Assignee:  c.String("assignee"),
		Creator:   c.String("creator"),
		Mentioned: c.String("mentioned"),
		Labels:    c.StringSlice("label"),
		Sort:      c.String("sort"),
		Direction: c.String("direction"),
		Since:     c.String("since"),
	}

	issues, err := client.GetIssues(repo, params)
	if err != nil {
		log.Fatal(err)
	}

	bold := false

	if len(issues) > 0 {
		PrintHeader()
		for _, issue := range issues {
			PrintIssue(issue, bold)
			bold = !bold
		}
	} else {
		fmt.Println("No issues found")
	}
}

var IssuesCommand cli.Command = cli.Command{
	Name:      "issues",
	ShortName: "is",
	Usage:     "List GitHub issues",
	Action:    Issues,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "milestone,m",
			Usage: "Milestone (none, *, or a milestone name)",
		},
		cli.StringFlag{
			Name:  "state,s",
			Usage: "Issue state (open, closed, all)",
			Value: "all",
		},
		cli.StringFlag{
			Name:  "assignee",
			Usage: "Assignee (none, *, or a username)",
		},
		cli.StringFlag{
			Name:  "creator",
			Usage: "Creator (must be a username)",
		},
		cli.StringFlag{
			Name:  "mentioned",
			Usage: "Mentioned (must be a username)",
		},
		cli.StringSliceFlag{
			Name:  "label,l",
			Value: &cli.StringSlice{},
			Usage: "Label filter to apply",
		},
		cli.StringFlag{
			Name:  "sort",
			Usage: "Sort by (created, updated, comments)",
			Value: "created",
		},
		cli.StringFlag{
			Name:  "direction",
			Usage: "Sort direction (asc, desc)",
			Value: "desc",
		},
		cli.StringFlag{
			Name:  "since",
			Usage: "Timestamp issue must be newer than (YYYY-MM-DDTHH:MM:SSZ format)",
		},
	},
}
