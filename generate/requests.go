package generate

import (
	"net/http"
	"net/url"

	"github.com/advanderveer/micros-parser/loader"
)

type Requests struct{}

func NewRequests() *Requests {
	return &Requests{}
}

func (r *Requests) Generate(host string, s *loader.Spec) ([]*http.Request, error) {
	reqs := []*http.Request{}

	for _, ep := range s.Endpoints {
		for _, c := range ep.Cases {

			l, err := url.Parse(host)
			if err != nil {
				return nil, err
			}

			l.Path = c.When.Path

			r, err := http.NewRequest(c.When.Method, l.String(), nil)
			if err != nil {
				return nil, err
			}

			reqs = append(reqs, r)
		}
	}

	return reqs, nil
}
