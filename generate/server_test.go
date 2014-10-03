package generate_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/advanderveer/micros/generate"
	"github.com/advanderveer/micros/loader"
)

func TestMockServer(t *testing.T) {
	rg := generate.NewTests()
	sets, err := rg.Generate(loader.FixNotesSpec(t))
	if err != nil {
		t.Fatal(err)
	}

	svr := generate.NewServer(sets)

	err = svr.Start()
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.Get(fmt.Sprintf("%s/notes", svr.URL()))
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != 201 {
		t.Fatal("Expected mock to succeed")
	}
}
