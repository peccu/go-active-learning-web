package expand_url_test

import (
	"testing"

	"github.com/codegangsta/cli"
	"github.com/syou6162/go-active-learning-web/lib/command"
)

func TestDoExpandUrl(t *testing.T) {
	app := cli.NewApp()
	app.Commands = command.Commands
	args := []string{
		"go-active-learning-web",
		"expand-url",
		"--input-filename=../../tech_input_example.txt",
	}

	if err := app.Run(args); err != nil {
		t.Error(err)
	}
}
