package loader_test

import (
	"strings"
	"testing"

	"github.com/advanderveer/micros/loader"
)

func TestCoditionsReading(t *testing.T) {
	bl := loader.NewBasic()

	s, err := bl.Load(strings.NewReader(ex1))
	if err != nil {
		t.Fatal(err)
	}

	if s.Endpoints[0].Name != "list_notes" {
		t.Fatal("Should have read name")
	}

	if s.Endpoints[0].Cases[0].When.Method != "GET" {
		t.Errorf("Unexpected method: %s", s.Endpoints[0].Cases[0].When.Method)
	}

	if s.Endpoints[0].Cases[0].When.Path != "/notes" {
		t.Errorf("Unexpected path: %s", s.Endpoints[0].Cases[0].When.Path)
	}
}

func TestExpectationsReading(t *testing.T) {
	bl := loader.NewBasic()

	s, err := bl.Load(strings.NewReader(ex1))
	if err != nil {
		t.Fatal(err)
	}

	if s.Endpoints[0].Cases[0].Then.StatusCode != 201 {
		t.Errorf("Expected statuscode not to be: %d", s.Endpoints[0].Cases[0].Then.StatusCode)
	}
}
