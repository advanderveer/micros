package command_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/advanderveer/micros/command"
)

func TestNotesPreAndEnv(t *testing.T) {
	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" || r.Method == "DELETE" {
			w.WriteHeader(200)
		} else {
			w.WriteHeader(201)
		}

	}))

	out := bytes.NewBuffer(nil)
	cmd := command.NewTest(out)

	//expect to output env data two times
	AssertCommand(t, cmd, []string{"--pre='env'", "--pre='env'", "--spec=../examples/notes.json", svr.URL}, `(?s).*PATH.*PATH.*`, out)

}
