package command_test

import (
	"bytes"
	"testing"

	"github.com/advanderveer/micros/command"
)

func TestNotesPreAndEnv(t *testing.T) {

	out := bytes.NewBuffer(nil)
	cmd := command.NewTest(out)

	//expect to output env data two times
	AssertCommand(t, cmd, []string{"--pre='env'", "--pre='env'", "--spec=../examples/notes.json"}, `(?s).*PATH.*PATH.*`, out)

}
