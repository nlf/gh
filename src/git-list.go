package main

import (
	"fmt"
	"log"

	"github.com/nlf/go-gh"
)

func PrintHeader() {
	fmt.Printf("\x1b[1m%-5s %-100.100s %-15.15s %-15.15s %-5s\x1b[0m\n", "ID", "Title", "Assignee", "Milestone", "Comments")
}

func PrintIssue(issue github.Issue, bold bool) {
	pullRequest := issue.PullRequest.URL != ""
	lineFormat := ""

	if bold {
		if pullRequest {
			lineFormat = "\x1b[1m%-5d [pr] %-95.95s %-15.15s %-15.15s %-5d\x1b[0m\n"
		} else {
			lineFormat = "\x1b[1m%-5d %-100.100s %-15.15s %-15.15s %-5d\x1b[0m\n"
		}
	} else {
		if pullRequest {
			lineFormat = "%-5d [pr] %-95.95s %-15.15s %-15.15s %-5d\n"
		} else {
			lineFormat = "%-5d %-100.100s %-15.15s %-15.15s %-5d\n"
		}
	}

	fmt.Printf(lineFormat, issue.Number, issue.Title, issue.Assignee.Login, issue.Milestone.Title, issue.Comments)
}

func main() {
	client := &github.Client{}
	client.LoadConfig()

	repo := github.GetRepo()

	issues, err := client.GetIssues(repo)
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
