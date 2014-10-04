package command

import (
	"io"
	"text/template"

	"github.com/codegangsta/cli"
)

var tmpl = `tested!`

type Test struct {
	*cmd
}

func NewTest(out io.Writer) *Test {
	return &Test{
		cmd: newCmd(out),
	}
}

func (c *Test) Name() string {
	return "test"
}

func (c *Test) Description() string {
	return "Test dependent services"
}

func (c *Test) Usage() string {
	return "Test dependent services."
}

func (c *Test) Flags() []cli.Flag {
	return []cli.Flag{
		cli.StringSliceFlag{Name: "pre", Value: &cli.StringSlice{}, Usage: ""},
	}
}

func (c *Test) Action() func(ctx *cli.Context) {
	return c.templated(c.Run)
}

func (c *Test) Run(ctx *cli.Context) (*template.Template, interface{}, error) {

	//load service spec

	//get dependencies

	//Test each dependencies

	//@todo implement

	return template.Must(template.New("test.success").Parse(tmpl)), nil, nil
}
