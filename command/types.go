package command

import (
	"fmt"
	"io"
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

type cmd struct {
	out io.Writer
}

func newCmd(out io.Writer) *cmd {
	if out == nil {
		out = os.Stdout
	}

	return &cmd{out}
}

func (c *cmd) Run(ctx *cli.Context) (*template.Template, interface{}, error) {
	return nil, nil, fmt.Errorf("Command '%s' is not yet implemented", ctx.Command.Name)
}

func (c *cmd) templated(fn func(c *cli.Context) (*template.Template, interface{}, error)) func(ctx *cli.Context) {
	return func(ctx *cli.Context) {
		t, data, err := fn(ctx)
		if err != nil {
			log.Fatal(err, ", Command: '", ctx.Command.Name, "' Args: ", ctx.Args())
		}

		err = t.Execute(c.out, data)
		if err != nil {
			log.Fatal("Template Error: ", err)
		}
	}
}
