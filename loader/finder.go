package loader

import (
	"fmt"
	"io"
	"os"
	"path"
)

type Finder struct {
	Entry string
}

func NewFinder(entry string) *Finder {
	return &Finder{entry}
}

func (f *Finder) Find(name string) (io.Reader, error) {
	dir := path.Dir(f.Entry)

	src, err := os.Open(path.Join(dir, fmt.Sprintf("%s.json", name)))
	if err != nil {
		return nil, fmt.Errorf("could not open %s, (entry: %s), %s", name, f.Entry, err)
	}

	return src, nil
}
