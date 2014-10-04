package generate

import (
	"net/http/httptest"
	"os"

	"github.com/zenazn/goji/web"

	"github.com/advanderveer/micros/loader"
)

// A mock server instance
type Server struct {
	sets []*TestSet
	mux  *web.Mux
	svr  *httptest.Server
}

func NewServer(sets []*TestSet) *Server {
	s := &Server{
		sets: sets,
		mux:  web.New(),
	}

	for _, set := range sets {
		s.mux.Handle(set.Pattern, set.Mock)
	}

	s.svr = httptest.NewUnstartedServer(s.mux)
	return s
}

func (s *Server) Start() error {
	s.svr.Start()
	return nil
}

func (s *Server) Stop() {
	s.svr.CloseClientConnections()
	s.svr.Close()
}

func (s *Server) URL() string {
	return s.svr.URL
}

// Creates mock server instances
type Factory struct {
	finder    *loader.Finder
	loader    *loader.Basic
	generator *Tests
	servers   map[string]*Server
}

func NewFactory(f *loader.Finder) *Factory {
	fac := &Factory{
		finder:  f,
		loader:  loader.NewBasic(),
		servers: map[string]*Server{},
	}

	fac.generator = NewTests(fac)
	return fac
}

func (f *Factory) Create(name string) (*Server, error) {

	//already have server with the name
	if svr, ok := f.servers[name]; ok {
		return svr, nil
	}

	//find spec source by name
	src, err := f.finder.Find(name)
	if err != nil {
		return nil, err
	}

	spec, err := f.loader.Load(src)
	if err != nil {
		return nil, err
	}

	//close if its a file
	if file, ok := src.(*os.File); ok {
		file.Close()
	}

	sets, err := f.generator.Generate(spec)
	if err != nil {
		return nil, err
	}

	f.servers[name] = NewServer(sets)
	return f.servers[name], nil
}
