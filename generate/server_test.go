package generate_test

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/advanderveer/micros/generate"
	"github.com/advanderveer/micros/loader"
)

func TestMockServer(t *testing.T) {
	f := loader.NewFinder("../examples/notes.json")
	fac := generate.NewFactory(f)
	rg := generate.NewTests(fac)
	sets, err := rg.Generate(loadSpec(t, f.Entry))
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

	if resp.StatusCode != 200 {
		t.Fatal("Expected mock to succeed")
	}
}

func TestServerFactory(t *testing.T) {

	f := loader.NewFinder("../examples/notes.json")
	fac := generate.NewFactory(f)

	svr1, err := fac.Create("notes")
	if err != nil {
		t.Fatal(err)
	}

	svr2, err := fac.Create("notes")
	if err != nil {
		t.Fatal(err)
	}

	if svr1 != svr2 {
		t.Fatal("Expected factory create to cache result")
	}
}
