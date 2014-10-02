package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"

	"github.com/advanderveer/micros-parser/command"
)

var version = "0.0.0-DEV"
var build = "unbuild"

func main() {
	app := cli.NewApp()
	app.Name = "micros"
	app.Usage = "micro-service test and development environment"
	app.Version = fmt.Sprintf("%s (%s)", version, build)

	//init micros commands
	cmds := []command.C{
		command.NewMock(),
	}

	//append to app
	for _, c := range cmds {
		app.Commands = append(app.Commands, cli.Command{
			Name:        c.Name(),
			Usage:       c.Usage(),
			Action:      c.Action(),
			Description: c.Description(),
			Flags:       c.Flags(),
		})
	}

	app.Run(os.Args)
}
