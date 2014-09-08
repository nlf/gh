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

	issues, err := client.GetIssues(repo, c.StringSlice("label"), c.String("milestone"))
	if err != nil {
		log.Fatal(err)
	}

	bold := false

	PrintHeader()
	if len(issues) > 0 {
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
			Usage: "Milestone filter to apply",
		},
		cli.StringSliceFlag{
			Name:  "label,l",
			Value: &cli.StringSlice{},
			Usage: "Label filter to apply",
		},
	},
}
