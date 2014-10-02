package generate

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/advanderveer/micros-parser/loader"
)

type TestSet struct {
	Request *http.Request
	Assert  AssertFunc
	Test    TestFunc
	Mock    http.HandlerFunc
}

type TestFunc func(c *http.Client) error
type AssertFunc func(resp *http.Response) error

type Tests struct{}

func NewTests() *Tests {
	return &Tests{}
}

func (tg *Tests) generateRequest(host string, c *loader.Case) (*http.Request, error) {

	l, err := url.Parse(host)
	if err != nil {
		return nil, err
	}

	// @todo specify more then just the path

	l.Path = c.When.Path

	return http.NewRequest(c.When.Method, l.String(), nil)
}

func (tg *Tests) generateAssert(c *loader.Case) (AssertFunc, error) {
	return func(resp *http.Response) error {

		// @todo make checks more sofisticated based on spec

		if c.Then.StatusCode != resp.StatusCode {
			return fmt.Errorf("Receiver status code %d, expected %d", c.Then.StatusCode, resp.StatusCode)
		}

		return nil
	}, nil
}

func (tg *Tests) generateMock(c *loader.Case) (http.HandlerFunc, error) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// @todo write mock handlers more sofisticaed according to spec

		w.WriteHeader(c.Then.StatusCode)

	}), nil
}

func (tg *Tests) generateTest(req *http.Request, a AssertFunc) (TestFunc, error) {
	return func(c *http.Client) error {

		resp, err := c.Do(req)
		if err != nil {
			return err
		}

		return a(resp)
	}, nil
}

func (tg *Tests) Generate(host string, s *loader.Spec) ([]*TestSet, error) {
	tests := []*TestSet{}

	for _, ep := range s.Endpoints {
		for _, c := range ep.Cases {
			r, err := tg.generateRequest(host, c)
			if err != nil {
				return nil, err
			}

			a, err := tg.generateAssert(c)
			if err != nil {
				return nil, err
			}

			t, err := tg.generateTest(r, a)
			if err != nil {
				return nil, err
			}

			m, err := tg.generateMock(c)
			if err != nil {
				return nil, err
			}

			tests = append(tests, &TestSet{r, a, t, m})
		}
	}

	return tests, nil
}
