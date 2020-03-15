package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func (receiver *Server) GorillaInit(addr string) {
	router := mux.NewRouter()
	router.HandleFunc("/", handleRedirect)
	router.HandleFunc("/upload", receiver.handleUpload())
	router.HandleFunc("/favicon.ico", receiver.handleFavicon())

	http.Handle("/", router)
	fmt.Println("Server is listening...")
	http.ListenAndServe(addr, nil)
}
