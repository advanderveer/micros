package generate_test

import (
	"os"
	"testing"

	"github.com/advanderveer/micros/loader"
)

func loadSpec(t *testing.T, path string) *loader.Spec {
	bl := loader.NewBasic()

	sfile, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer sfile.Close()

	s, err := bl.Load(sfile)
	if err != nil {
		t.Fatal(err)
	}

	return s
}
