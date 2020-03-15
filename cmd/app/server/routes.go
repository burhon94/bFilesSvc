package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func (receiver *Server) GorillaInitRoutes(addr string) {
	router := mux.NewRouter()
	router.HandleFunc("/", indexHandler)

	http.Handle("/", router)
	fmt.Println("Server is listening...")
	http.ListenAndServe(addr, nil)
}
