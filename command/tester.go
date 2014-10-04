package command

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/codegangsta/cli"

	"github.com/advanderveer/micros/generate"
	"github.com/advanderveer/micros/loader"
)

var tmpl_test = `tested!`

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
	return "Test a service"
}

func (c *Test) Usage() string {
	return "Test a micro service."
}

func (c *Test) Flags() []cli.Flag {
	return []cli.Flag{
		cli.StringSliceFlag{Name: "pre, p", Value: &cli.StringSlice{}, Usage: "Execute command after the mocks where setupt and before the tests are run"},
		cli.StringFlag{Name: "spec, s", Value: "", Usage: "Provide the path to a local spec"},
		cli.StringFlag{Name: "runner", Value: "sh -c {{.}}", Usage: "Shell wrapper that runs each command."},
	}
}

func (c *Test) Action() func(ctx *cli.Context) {
	return c.templated(c.Run)
}

func (c *Test) Run(ctx *cli.Context) (*template.Template, interface{}, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, nil, err
	}

	target := ctx.Args().First()
	if target == "" {
		return nil, nil, NoTargetError
	}

	spath := strings.TrimSpace(ctx.String("spec"))
	if spath == "" {
		return nil, nil, NoSpecPathError
	}

	spec, err := c.loadSpec(filepath.Join(wd, spath))
	if err != nil {
		return nil, nil, err
	}

	//@todo launch dependency mocks

	//parse runner template
	r := ctx.String("runner")
	for _, pre := range ctx.StringSlice("pre") {
		cmdparts := []string{}

		//so we compile each part of the runner seperately so
		//argument seperation is maintained correctly when input is complex
		for _, p := range strings.Split(r, " ") {

			rt, err := template.New("runner").Parse(p)
			if err != nil {
				return nil, nil, fmt.Errorf("Error while parsing runner template (%s), part: (%s): %s", r, p, err)
			}

			b := bytes.NewBuffer(nil)
			err = rt.Execute(b, pre)
			if err != nil {
				return nil, nil, fmt.Errorf("Error while executing runner template (%s, part: %s): %s", r, pre, err)
			}

			cmdparts = append(cmdparts, b.String())
		}

		//create command to be run
		cmd := exec.Command(cmdparts[0], cmdparts[1:]...)
		cmd.Stdout = c.out
		cmd.Stderr = c.out

		//@todo set env variables

		//run the command
		err := cmd.Run()
		if err != nil {
			return nil, nil, fmt.Errorf("Failed: %s", err)
		}

	}

	//server factory
	f := loader.NewFinder(spath)
	fac := generate.NewFactory(f)

	//generate test sets
	tgen := generate.NewTests(fac)
	sets, err := tgen.Generate(spec)
	if err != nil {
		return nil, nil, err
	}

	//runs every case
	for _, set := range sets {
		err := set.Test(target, http.DefaultClient)
		if err != nil {
			return nil, nil, err
		}
	}

	return template.Must(template.New("test.success").Parse(tmpl_test)), nil, nil
}
