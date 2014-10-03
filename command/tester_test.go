package command_test

import (
	"flag"
	"testing"

	"github.com/codegangsta/cli"

	"github.com/advanderveer/micros/command"
)

func TestNotesPreAndEnv(t *testing.T) {
	app := cli.NewApp()

	set := flag.NewFlagSet("x", 0)
	set.Parse([]string{"--pre=aaaa"})

	ctx := cli.NewContext(app, set, set)

	cmd := command.NewTest()

	AssertOutput(t, ctx, `.*tested.*`, cmd.Run)
}
