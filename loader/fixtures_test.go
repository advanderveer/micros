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
						"status_code": 201
					}
				}
			]
		}
	]
}`

func FixNotesSpec(t *testing.T) *loader.Spec {
	bl := loader.NewBasic()

	s, err := bl.Load(strings.NewReader(ex1))
	if err != nil {
		t.Fatal(err)
	}

	return s
}
