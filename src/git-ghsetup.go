package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"code.google.com/p/gopass"
	"github.com/nlf/go-gh/lib"
)

func Prompt(prompt string) string {
	fmt.Printf(prompt)
	stdin := bufio.NewReader(os.Stdin)
	line, _, err := stdin.ReadLine()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return string(line)
}

func main() {
	user := Prompt("Username: ")

	password, err := gopass.GetPass("Password: ")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	has2fa := Prompt("Use two-factor authentication? (y/n) ")
	has2fa = strings.ToLower(has2fa)
	token := ""

	if has2fa == "y" || has2fa == "yes" {
		token = Prompt("Token: ")
	}

	baseUrl := Prompt("Base URL (https://api.github.com): ")
	if baseUrl == "" {
		baseUrl = "https://api.github.com"
	}

	generatedToken, err := github.CreateToken(baseUrl, user, password, token)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(generatedToken)
	os.Exit(0)
}
