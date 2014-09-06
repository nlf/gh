package github

import (
	"log"
	"os/exec"
	"strings"
)

func GetRepo() string {
	remotes, err := exec.Command("git", "remote", "-v").Output()
	if err != nil {
		log.Fatal(err)
	}

	remotesString := strings.Trim(string(remotes), " \n")

	lines := strings.Split(remotesString, "\n")
	line := ""

	for _, line = range lines {
		if strings.HasPrefix(line, "origin") && strings.HasSuffix(line, "(push)") {
			break
		}
	}

	line = line[strings.Index(line, ":")+1:]
	line = line[:strings.Index(line, " (push)")]
	lineSegments := strings.Split(line, "/")
	repo := strings.Join(lineSegments[len(lineSegments)-2:], "/")
	repo = strings.TrimSuffix(repo, ".git")

	return repo
}

func GetRoot() string {
	dir, err := exec.Command("git", "rev-parse", "--show-toplevel").Output()
	if err != nil {
		log.Fatal(err)
	}

	return strings.Trim(string(dir), " \n")
}

func GetBranch() string {
	branch, err := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD").Output()
	if err != nil {
		log.Fatal(err)
	}

	return strings.Trim(string(branch), " \n")
}
