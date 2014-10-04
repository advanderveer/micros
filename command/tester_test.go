package command_test

import (
	"bytes"
	"testing"

	"github.com/advanderveer/micros/command"
)

func TestNotesPreAndEnv(t *testing.T) {

	out := bytes.NewBuffer(nil)
	cmd := command.NewTest(out)

	AssertCommand(t, cmd, []string{"--pre=a", "--pre=b"}, `.*tested.*`, out)
}
