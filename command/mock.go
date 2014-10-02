package command

import (
	"github.com/codegangsta/cli"
)

type Mock struct {
	*cmd
}

func NewMock() *Mock {
	return &Mock{
		cmd: &cmd{},
	}
}

func (cmd *Mock) Name() string {
	return "mock"
}

func (cmd *Mock) Description() string {
	return "Mock dependent services"
}

func (cmd *Mock) Usage() string {
	return "Mock dependent services."
}

func (cmd *Mock) Flags() []cli.Flag {
	return []cli.Flag{}
}
