package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/syou6162/go-active-learning-web/lib/command"
)

func main() {
	app := cli.NewApp()
	app.Name = "go-active-learning-web"
	app.Commands = command.Commands

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
