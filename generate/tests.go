package generate

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/zenazn/goji/web"

	"github.com/advanderveer/micros/loader"
)

//One set per test case
type TestSet struct {
	Request           *http.Request
	Assert            AssertFunc
	Test              TestFunc
	Mock              web.HandlerFunc
	Pattern           web.Pattern
	DependencyServers []*Server
}

type TestFunc func(host string, c *http.Client) error
type AssertFunc func(resp *http.Response) error

//Generated pattern for matching mock requests to the correct mock handler
type MockPattern struct {
	Prototype *http.Request
}

func (p *MockPattern) Prefix() string {
	return "/"
}

// Match the received request to the handler by examining the prototype request
func (p *MockPattern) Match(r *http.Request, c *web.C) bool {

	//@todo compare incoming request with prototype request

	if r.URL.Path != p.Prototype.URL.Path {
		return false
	}

	if r.Method != p.Prototype.Method {
		return false
	}

	return true
}

func (p *MockPattern) Run(r *http.Request, c *web.C) {}

//A Test set generator
type Tests struct {
	factory *Factory

	IgnoreDependencyChecks bool
}

func NewTests(f *Factory) *Tests {
	return &Tests{
		factory: f,
	}
}

// Generates an request to be send to the subject based on the case provided by the spec
func (tg *Tests) generateRequest(c *loader.Case) (*http.Request, error) {

	l, err := url.Parse("/")
	if err != nil {
		return nil, err
	}

	// @todo start recording

	// @todo specify more then just the path

	l.Path = c.When.Path

	return http.NewRequest(c.When.Method, l.String(), nil)
}

// Generate an assertion function that checks the response returned by the subject
func (tg *Tests) generateAssert(c *loader.Case) (AssertFunc, error) {
	return func(resp *http.Response) error {

		// @todo make checks more sofisticated based on spec

		if c.Then.StatusCode != resp.StatusCode {
			return fmt.Errorf("Receiver status code %d, expected %d", c.Then.StatusCode, resp.StatusCode)
		}

		return nil
	}, nil
}

func (tg *Tests) generateDependencyServers(c *loader.Case) ([]*Server, error) {
	svrs := []*Server{}

	for _, dep := range c.While {
		svr, err := tg.factory.Create(dep.Service)
		if err != nil {
			return nil, err
		}

		svrs = append(svrs, svr)
	}

	return svrs, nil
}

func (tg *Tests) generatePattern(r *http.Request) (*MockPattern, error) {
	return &MockPattern{r}, nil
}

// Generate the http handler function that writes the expected response based on the specification
func (tg *Tests) generateMock(c *loader.Case, svrs []*Server) (web.HandlerFunc, error) {
	return web.HandlerFunc(func(ctx web.C, w http.ResponseWriter, r *http.Request) {

		// @todo write mock handlers more sofisticaed according to spec

		w.WriteHeader(c.Then.StatusCode)

		// mock should also call the dependencies
		for _, dep := range svrs {
			loc := dep.URL()
			if loc == "" {
				dep.Start()
				loc = dep.URL()
			}

			_, err := http.Get(dep.URL())
			if err != nil {
				http.Error(w, fmt.Sprintf("Mock couldn't GET dep service @ %s", dep.URL()), http.StatusInternalServerError)
			}
		}

	}), nil
}

func (tg *Tests) generateTest(req *http.Request, a AssertFunc, svrs []*Server) (TestFunc, error) {
	return func(host string, c *http.Client) error {

		//parse overwrite host url
		h, err := url.Parse(host)
		if err != nil {
			return err
		}

		//overwrite generated with test specific host/scheme
		req.URL.Host = h.Host
		req.URL.Scheme = h.Scheme

		//start recording on each the dep servers
		for _, svr := range svrs {
			svr.Record()
			defer svr.Rewind()
		}

		//do the actual request
		resp, err := c.Do(req)
		if err != nil {
			return err
		}

		//@todo figure out how to give tested service time (timeout) to send request to dependency

		//check recording
		//@todo not only check if the svr was reached
		//@todo eleborate error message
		for _, svr := range svrs {
			if !tg.IgnoreDependencyChecks && svr.Tape != nil && len(svr.Tape) < 1 {
				return fmt.Errorf("Dependency did not receive the expected request")
			}
		}

		//and assert resp
		return a(resp)
	}, nil
}

func (tg *Tests) Generate(s *loader.Spec) ([]*TestSet, error) {
	tests := []*TestSet{}

	for _, ep := range s.Endpoints {
		for _, c := range ep.Cases {

			r, err := tg.generateRequest(c)
			if err != nil {
				return nil, err
			}

			a, err := tg.generateAssert(c)
			if err != nil {
				return nil, err
			}

			svrs, err := tg.generateDependencyServers(c)
			if err != nil {
				return nil, err
			}

			t, err := tg.generateTest(r, a, svrs)
			if err != nil {
				return nil, err
			}

			m, err := tg.generateMock(c, svrs)
			if err != nil {
				return nil, err
			}

			p, err := tg.generatePattern(r)
			if err != nil {
				return nil, err
			}

			tests = append(tests, &TestSet{r, a, t, m, p, svrs})
		}
	}

	return tests, nil
}
