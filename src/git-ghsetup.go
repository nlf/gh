package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"code.google.com/p/gopass"
	"github.com/nlf/go-gh"
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

func main() {
	user := Prompt("Username: ")

	password, err := gopass.GetPass("Password: ")
	if err != nil {
		log.Fatal(err)
	}

	has2fa := Prompt("Use two-factor authentication? (y/n) ")
	has2fa = strings.ToLower(has2fa)
	token := ""

	if has2fa == "y" || has2fa == "yes" {
		token = Prompt("Token: ")
	}

	baseUrl := Prompt("API URL (https://api.github.com): ")
	if baseUrl == "" {
		baseUrl = "https://api.github.com"
	}

	generatedToken, err := github.CreateToken(baseUrl, user, password, token)
	if err != nil {
		log.Fatal(err)
	}

	client := &github.Client{BaseURL: baseUrl, Token: generatedToken.Token}
	client.SaveConfig()

	os.Exit(0)
}
