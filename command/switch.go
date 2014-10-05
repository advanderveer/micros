package command

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/advanderveer/micros/loader"
	"github.com/codegangsta/cli"
)

var tmpl_switch = `Switched to {{.}}`

// workspace configuration
type Workspace struct {
	Entry string `json:"entry"`
}

// switches between services in the workspace
type Switch struct {
	*cmd
}

func NewSwitch(out io.Writer) *Switch {
	return &Switch{
		cmd: newCmd(out),
	}
}

func (c *Switch) Name() string {
	return "switch"
}

func (c *Switch) Description() string {
	return "Switch between micro-services."
}

func (c *Switch) Usage() string {
	return "Switch between micro-services."
}

func (c *Switch) Flags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{Name: "dir, d", Value: "", Usage: fmt.Sprintf("Specify another directory then the current working dir.")},
	}
}

func (c *Switch) Action() func(ctx *cli.Context) {
	return c.templated(c.Run)
}

func (c *Switch) Run(ctx *cli.Context) (*template.Template, interface{}, error) {
	var err error

	//switch to
	to := ctx.Args().First()
	if to == "" {
		return nil, nil, NoServiceNameError
	}

	dir := strings.TrimSpace(ctx.String("dir"))
	if dir == "" {
		dir, err = os.Getwd()
		if err != nil {
			return nil, nil, err
		}

	}

	f, err := os.OpenFile(filepath.Join(dir, "micros.json"), os.O_RDWR, 0777)
	if err != nil {
		return nil, nil, err
	}
	defer f.Close()

	//load workspace configuration
	space := Workspace{}
	dec := json.NewDecoder(f)
	err = dec.Decode(&space)
	if err != nil {
		return nil, nil, err
	}

	//find service spec by name
	finder := loader.NewFinder(f.Name())
	space.Entry, err = finder.FindPath(to)
	if err != nil {
		return nil, nil, err
	}

	//trunc and write workspace file
	err = f.Truncate(0)
	if err != nil {
		return nil, nil, err
	}

	//this is not very safe way of writing
	enc := json.NewEncoder(f)
	err = enc.Encode(space)
	if err != nil {
		return nil, nil, err
	}

	return template.Must(template.New("switch.success").Parse(tmpl_switch)), space.Entry, nil
}
