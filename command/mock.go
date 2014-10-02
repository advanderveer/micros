package command

import (
	"text/template"

	"github.com/codegangsta/cli"
)

var tmpl = `mocked!`

type Mock struct {
	*cmd
}

func NewMock() *Mock {
	return &Mock{
		cmd: &cmd{},
	}
}

func (c *Mock) Name() string {
	return "mock"
}

func (c *Mock) Description() string {
	return "Mock dependent services"
}

func (c *Mock) Usage() string {
	return "Mock dependent services."
}

func (c *Mock) Flags() []cli.Flag {
	return []cli.Flag{}
}

func (c *Mock) Run(ctx *cli.Context) (*template.Template, interface{}, error) {

	//load service spec

	//get dependencies

	//mock each dependencies

	//@todo implement

	return template.Must(template.New("mock.success").Parse(tmpl)), nil, nil
}
