package generate_test

import (
	"net/http"
	"testing"

	"github.com/advanderveer/micros/generate"
	"github.com/advanderveer/micros/loader"
)

func TestMockServer(t *testing.T) {
	rg := generate.NewTests()
	sets, err := rg.Generate("http://locahost", loader.FixNotesSpec(t))
	if err != nil {
		t.Fatal(err)
	}

	svr := generate.NewServer(sets)

	err = svr.Start()
	if err != nil {
		t.Fatal(err)
	}

	resp, err := http.Get(svr.URL())
	if err != nil {
		t.Fatal(err)
	}

	_ = resp
	//@assert mocked resp

}
