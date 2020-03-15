package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Server struct {
	gRouter *mux.Router
}

func NewServer(gRouter *mux.Router) *Server {
	return &Server{gRouter: gRouter}
}

func (receiver *Server) Start(addr string)  {
	receiver.GorillaInit(addr)
}

func (receiver *Server) ServeHTTP(w http.ResponseWriter, r *http.Request)  {
	receiver.gRouter.ServeHTTP(w, r)
}

