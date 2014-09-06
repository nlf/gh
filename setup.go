package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"code.google.com/p/gopass"
	"github.com/codegangsta/cli"
	"github.com/nlf/gh/github"
)

func Prompt(prompt string) string {
	fmt.Printf(prompt)
	stdin := bufio.NewReader(os.Stdin)
	line, _, err := stdin.ReadLine()

	if err != nil {
		log.Fatal(err)
	}

	return string(line)
}

func Setup(c *cli.Context) {
	user := Prompt("Username: ")

	password, err := gopass.GetPass("Password: ")
	if err != nil {
		log.Fatal(err)
	}

	token := ""
	if c.Bool("two-factor") {
		token = Prompt("Token: ")
	}

	baseUrl := c.String("url")

	generatedToken, err := github.CreateToken(baseUrl, user, password, token)
	if err != nil {
		log.Fatal(err)
	}

	client := &github.Client{BaseURL: baseUrl, Token: generatedToken.Token}
	client.SaveConfig()

	os.Exit(0)
}

var SetupCommand cli.Command = cli.Command{
	Name:   "setup",
	Usage:  "create a configuration for gh",
	Action: Setup,
	Flags: []cli.Flag{
		cli.BoolFlag{
			Name:  "two-factor,2fa",
			Usage: "Enable two-factor authentication",
		},
		cli.StringFlag{
			Name:  "url",
			Value: "https://api.github.com",
			Usage: "URL to use for the GitHub API",
		},
	},
}
