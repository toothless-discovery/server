package server

import (
	"log"
	"net/http"

	"github.com/toothless-discovery/server/services"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

// Server type for toothless-discovery
type Server struct {
	server *rpc.Server
	router *mux.Router
}

// New toothless-discovery server
func New() *Server {
	server := rpc.NewServer()
	server.RegisterCodec(json.NewCodec(), "application/json")
	server.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")

	router := mux.NewRouter()
	router.Handle("/rpc", server)

	ret := &Server{
		server: server,
		router: router,
	}

	err := ret.RegisterService(new(services.Discovery), "Discovery")
	if err != nil {
		log.Fatalf("Unable to register discovery service.")
	}

	return ret
}

//RegisterService add new json-rpc handler to server
func (s *Server) RegisterService(receiver interface{}, name string) error {
	return s.server.RegisterService(receiver, name)
}

// Handler for http server
func (s *Server) Handler() http.Handler {
	return s.router
}
