package generate_test

import (
	"testing"

	"github.com/advanderveer/micros-parser/generate"
	"github.com/advanderveer/micros-parser/loader"
)

func TestGenerate(t *testing.T) {
	rg := generate.NewRequests()

	reqs, err := rg.Generate("http://locahost/bogus?foo=bar", loader.FixNotesSpec(t))
	if err != nil {
		t.Fatal(err)
	}

	if len(reqs) != 1 {
		t.Fatal("Expected 1 request to be generated")
	}

	if reqs[0].URL.String() != "http://locahost/notes?foo=bar" {
		t.Fatalf("Expected unexpected request url: %s", reqs[0].URL.String())
	}

}
