package generate

import (
	"net/http/httptest"

	"github.com/zenazn/goji/web"
)

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

	//@tood map sets to mux

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
