package loader

import (
	"strings"
	"testing"
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

func FixNotesSpec(t *testing.T) *Spec {
	bl := NewBasic()

	s, err := bl.Load(strings.NewReader(ex1))
	if err != nil {
		t.Fatal(err)
	}

	return s
}
