package command_test

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"

	"github.com/advanderveer/micros/command"
)

var ex1 = `{}`

func TestSwitch(t *testing.T) {
	dir, err := ioutil.TempDir("", "t_micros_switch")
	if err != nil {
		t.Fatal(err)
	}

	err = ioutil.WriteFile(filepath.Join(dir, "micros.json"), []byte(ex1), 0777)
	if err != nil {
		t.Fatal(err)
	}

	out := bytes.NewBuffer(nil)
	cmd := command.NewSwitch(out)

	//expect to output env data two times
	AssertCommand(t, cmd, []string{"-d=" + dir, "notes"}, `(?s).*Switched.*notes.json.*`, out)

	data, err := ioutil.ReadFile(filepath.Join(dir, "micros.json"))
	if err != nil {
		t.Fatal(err)
	}

	if strings.Contains(string(data), "{}") {
		t.Fatal("Expected workspace file to be rewritten")
	}
}
