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
		http.Error(w, "test", 404)
	}))

	defer ts.Close()

	rg := generate.NewTests()

	sets, err := rg.Generate(ts.URL, loader.FixNotesSpec(t))
	if err != nil {
		t.Fatal(err)
	}

	err = sets[0].Test(http.DefaultClient)
	if err == nil {
		t.Fatal("Expected test to fail")
	}
}

func TestMocking(t *testing.T) {
	rg := generate.NewTests()
	sets, err := rg.Generate("http://locahost", loader.FixNotesSpec(t))
	if err != nil {
		t.Fatal(err)
	}

	ts := httptest.NewServer(sets[0].Mock)

	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != 201 {
		t.Fatal("Expected mock to simulate correctly")
	}

	if sets[0].Assert(resp) != nil {
		t.Fatal("Asserting response with own assert should always succeed")
	}

}
