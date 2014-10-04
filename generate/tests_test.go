package generate_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/advanderveer/micros/generate"
	"github.com/advanderveer/micros/loader"
)

func TestTestGenerate(t *testing.T) {
	f := loader.NewFinder("../examples/notes.json")
	fac := generate.NewFactory(f)
	rg := generate.NewTests(fac)

	tests, err := rg.Generate(loadSpec(t, "../examples/notes.json"))
	if err != nil {
		t.Fatal(err)
	}

	if len(tests) < 1 {
		t.Fatal("Expected more then zero tests to be generated")
	}
}

func TestTestExecution(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "test", 200)
	}))

	defer ts.Close()

	f := loader.NewFinder("../examples/notes.json")
	fac := generate.NewFactory(f)
	rg := generate.NewTests(fac)

	sets, err := rg.Generate(loadSpec(t, "../examples/notes.json"))
	if err != nil {
		t.Fatal(err)
	}

	err = sets[0].Test(ts.URL, http.DefaultClient)
	if err != nil {
		t.Fatalf("Expected test to succeed, but got %s", err)
	}
}

func TestMocking(t *testing.T) {
	f := loader.NewFinder("../examples/notes.json")
	fac := generate.NewFactory(f)
	rg := generate.NewTests(fac)
	sets, err := rg.Generate(loadSpec(t, "../examples/notes.json"))
	if err != nil {
		t.Fatal(err)
	}

	ts := httptest.NewServer(sets[0].Mock)

	resp, err := http.Get(ts.URL)
	if err != nil {
		t.Error(err)
	}

	if resp.StatusCode != 200 {
		t.Fatal("Expected mock to simulate correctly")
	}

	if sets[0].Assert(resp) != nil {
		t.Fatal("Asserting response with own assert should always succeed")
	}

}
