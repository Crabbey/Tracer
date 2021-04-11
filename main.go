package main

import (
	"fmt"
	"github.com/gravitational/version"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {

	app := cli.NewApp()
	app.Name = "tracer"
	app.Usage = "tracer"
	app.Version = version.Get().Version + " (Commit " + version.Get().GitCommit + ")"
	app.EnableBashCompletion = true
	app.CommandNotFound = func(context *cli.Context, cmd string) {
		fmt.Printf("ERROR: Unknown command '%s'\n", cmd)
	}

	app.Action = func(context *cli.Context) error {
		tracer, err := NewTracer(context)
		if err != nil {
			cli.Exit(err, 1)
		}
		return tracer.Run()
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}