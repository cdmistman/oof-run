package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{}
	app.Authors = []*cli.Author{
		{
			Name:  "Colton D",
			Email: "colton@donn.io",
		},
		{
			Name:  "Chase C.",
			Email: "compton.cm@gmail.com",
		},
	}

	if err := app.Run(os.Args); err != nil {
		// todo
	}
	os.Exit(0)
}
