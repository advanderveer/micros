package command

import (
	"fmt"
	"log"
	"os"
	"text/template"

	"github.com/codegangsta/cli"
)

type C interface {
	Name() string
	Description() string
	Usage() string
	Run(c *cli.Context) (*template.Template, interface{}, error)
	Action() func(ctx *cli.Context)
	Flags() []cli.Flag
}

type cmd struct{}

func (c *cmd) Run(ctx *cli.Context) (*template.Template, interface{}, error) {
	return nil, nil, fmt.Errorf("Command '%s' is not yet implemented", ctx.Command.Name)
}

func (c *cmd) Action() func(ctx *cli.Context) {
	return func(ctx *cli.Context) {
		t, data, err := c.Run(ctx)
		if err != nil {
			log.Fatal(err, ", Command: '", ctx.Command.Name, "' Args: ", ctx.Args())
		}

		err = t.Execute(os.Stdout, data)
		if err != nil {
			log.Fatal("Template Error: ", err)
		}

		//end with newline
		os.Stdout.WriteString("\n")
	}
}
