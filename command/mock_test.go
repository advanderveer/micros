package command_test

import (
	"flag"
	"testing"

	"github.com/codegangsta/cli"

	"github.com/advanderveer/micros/command"
)

func TestNotesMocking(t *testing.T) {
	set := flag.NewFlagSet("test", 0)
	ctx := cli.NewContext(nil, set, nil)

	cmd := command.NewMock()

	AssertOutput(t, ctx, `.*mocked.*`, cmd.Run)
}
