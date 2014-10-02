package generate_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/advanderveer/micros-parser/generate"
	"github.com/advanderveer/micros-parser/loader"
)

func TestTestGenerate(t *testing.T) {
	rg := generate.NewTests()

	tests, err := rg.Generate("http://locahost/bogus?foo=bar", loader.FixNotesSpec(t))
	if err != nil {
		t.Fatal(err)
	}

	if len(tests) != 1 {
		t.Fatal("Expected 1 test to be generated")
	}

}

func TestTestExecution(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "test", 201)
	}))

	defer ts.Close()

	rg := generate.NewTests()

	tests, err := rg.Generate(ts.URL, loader.FixNotesSpec(t))
	if err != nil {
		t.Fatal(err)
	}

	err = tests[0](http.DefaultClient)
	if err == nil {
		t.Fatal("Expected test to fail")
	}
}
