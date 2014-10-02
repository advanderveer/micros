package generate

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/advanderveer/micros-parser/loader"
)

type TestFunc func(c *http.Client) error
type AssertFunc func(resp *http.Response) error

type Tests struct{}

func NewTests() *Tests {
	return &Tests{}
}

func (t *Tests) generateRequest(host string, c *loader.Case) (*http.Request, error) {

	l, err := url.Parse(host)
	if err != nil {
		return nil, err
	}

	l.Path = c.When.Path

	return http.NewRequest(c.When.Method, l.String(), nil)
}

func (t *Tests) generateAssert(c *loader.Case) (AssertFunc, error) {
	return func(resp *http.Response) error {

		//check status code mismatch
		if c.Then.StatusCode != resp.StatusCode {
			return fmt.Errorf("Receiver status code %d, expected %d", c.Then.StatusCode, resp.StatusCode)
		}

		return nil
	}, nil
}

func (t *Tests) generateTest(req *http.Request, a AssertFunc) (TestFunc, error) {
	return func(c *http.Client) error {

		resp, err := c.Do(req)
		if err != nil {
			return err
		}

		return a(resp)
	}, nil
}

func (t *Tests) Generate(host string, s *loader.Spec) ([]TestFunc, error) {
	tests := []TestFunc{}

	for _, ep := range s.Endpoints {
		for _, c := range ep.Cases {
			r, err := t.generateRequest(host, c)
			if err != nil {
				return nil, err
			}

			a, err := t.generateAssert(c)
			if err != nil {
				return nil, err
			}

			t, err := t.generateTest(r, a)
			if err != nil {
				return nil, err
			}

			tests = append(tests, t)
		}
	}

	return tests, nil
}
