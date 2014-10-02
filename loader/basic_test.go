package loader_test

import (
	"strings"
	"testing"

	"github.com/advanderveer/micros/loader"
)

var ex1 = `{
	"endpoints": [
		{
			"name": "list_notes",
			"cases": [
				{
					"when": {
						"method": "GET",
						"path": "/notes"
					},
					"then": {
						"status_code": 200
					}
				}
			]
		}
	]
}`

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

	if s.Endpoints[0].Cases[0].Then.StatusCode != 200 {
		t.Errorf("Expected statuscode to be: %d", s.Endpoints[0].Cases[0].Then.StatusCode)
	}
}
