package command_test

import (
	"bytes"
	"regexp"
	"testing"
	"text/template"

	"github.com/codegangsta/cli"
)

func AssertOutput(t *testing.T, ctx *cli.Context, pattern string, fn func(c *cli.Context) (*template.Template, interface{}, error)) {
	var out string

	tmpl, data, err := fn(ctx)
	if err != nil {
		t.Error(err)
		out = err.Error()
	} else {

		buff := bytes.NewBuffer(nil)
		err = tmpl.Execute(buff, data)
		if err != nil {
			t.Error(err)
		}

		out = buff.String()
	}

	m, err := regexp.MatchString(pattern, out)
	if err != nil {
		t.Fatal(err)
	}

	if !m {
		t.Errorf("Out didn't match expected pattern /%s/, received: %s", pattern, out)
	}
}
