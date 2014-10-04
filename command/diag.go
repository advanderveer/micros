package command

import (
	// "bytes"
	// "fmt"
	"io"
	"net/http"
	"os"
	// "os/exec"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/codegangsta/cli"

	"github.com/advanderveer/micros/generate"
	// "github.com/advanderveer/micros/loader"
)

var tmpl_diag = `diagnosed!`

type Diag struct {
	*cmd
}

func NewDiag(out io.Writer) *Diag {
	return &Diag{
		cmd: newCmd(out),
	}
}

func (c *Diag) Name() string {
	return "diag"
}

func (c *Diag) Description() string {
	return "Diagnose current service spec"
}

func (c *Diag) Usage() string {
	return "Diagnose current service spec."
}

func (c *Diag) Flags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{Name: "spec, s", Value: "", Usage: "Provide the path to a local spec"},
	}
}

func (c *Diag) Action() func(ctx *cli.Context) {
	return c.templated(c.Run)
}

func (c *Diag) Run(ctx *cli.Context) (*template.Template, interface{}, error) {
	wd, err := os.Getwd()
	if err != nil {
		return nil, nil, err
	}

	spath := strings.TrimSpace(ctx.String("spec"))
	if spath == "" {
		return nil, nil, NoSpecPathError
	}

	spec, err := c.loadSpec(filepath.Join(wd, spath))
	if err != nil {
		return nil, nil, err
	}

	//generate test sets
	tgen := generate.NewTests()
	sets, err := tgen.Generate(spec)
	if err != nil {
		return nil, nil, err
	}

	//mock the service itself
	svr := generate.NewServer(sets)
	svr.Start()
	defer svr.Stop()

	//runs tests on the mocked service
	for _, set := range sets {
		err := set.Test(svr.URL(), http.DefaultClient)
		if err != nil {
			return nil, nil, err
		}
	}

	return template.Must(template.New("diag.success").Parse(tmpl_diag)), nil, nil
}
