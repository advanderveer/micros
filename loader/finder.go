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

func (f *Finder) FindPath(name string) (string, error) {
	return path.Join(path.Dir(f.Entry), fmt.Sprintf("%s.json", name)), nil
}

func (f *Finder) Find(name string) (io.Reader, error) {
	path, err := f.FindPath(name)
	if err != nil {
		return nil, fmt.Errorf("could not open %s, (entry: %s), %s", name, f.Entry, err)
	}

	src, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("could not open %s, (entry: %s), %s", name, f.Entry, err)
	}

	return src, nil
}
