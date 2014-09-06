package main

import (
	"os"

	"github.com/codegangsta/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "gh"
	app.Usage = "GitHub command line tools"
	app.Commands = []cli.Command{SetupCommand, IssueCommand}
	app.Run(os.Args)
}
