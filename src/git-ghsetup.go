package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/nlf/go-gh/lib"

	"code.google.com/p/gopass"
)

func Prompt(prompt string) (string, error) {
	fmt.Printf(prompt)
	stdin := bufio.NewReader(os.Stdin)
	line, _, err := stdin.ReadLine()

	return string(line), err
}

func ExitWithError(err error) {
	fmt.Println(err)
	os.Exit(1)
}

func main() {
	user, err := Prompt("Username: ")
	if err != nil {
		ExitWithError(err)
	}

	password, err := gopass.GetPass("Password: ")
	if err != nil {
		ExitWithError(err)
	}

	has2fa, err := Prompt("Use two-factor authentication? (y/n) ")
	if err != nil {
		ExitWithError(err)
	}

	has2fa = strings.ToLower(has2fa)
	token := ""

	if has2fa == "y" || has2fa == "yes" {
		token, err = Prompt("Token: ")
		if err != nil {
			ExitWithError(err)
		}
	}

	generatedToken, err := github.CreateToken(user, password, token)
	if err != nil {
		ExitWithError(err)
	}

	fmt.Println(generatedToken)
	os.Exit(0)

	// client := &http.Client{}
	// req, err := http.NewRequest("GET", "https://api.github.com/users/nlf", nil)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// req.SetBasicAuth(user, passwd)
	// req.Header.Add("X-GitHub-OTP", token)
	//
	// resp, err := client.Do(req)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// defer resp.Body.Close()
	//
	// ghUser := &githubUser{}
	// json.NewDecoder(resp.Body).Decode(&ghUser)
	//
	// fmt.Println(ghUser)
}
