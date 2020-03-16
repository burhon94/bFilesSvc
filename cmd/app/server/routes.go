package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func (receiver *Server) GorillaInit(addr string) {
	router := mux.NewRouter()
	router.HandleFunc("/api/health", receiver.handlerRespHealth()).Methods("get")

	//post files
	router.HandleFunc("/api/files/", receiver.handleUploading()).Methods("post")
	//get file
	router.PathPrefix("/api/files/").HandlerFunc(receiver.handleGetFile()).Methods("get")

	http.Handle("/", router)
	fmt.Println("Server is listening...")
	if http.ListenAndServe(addr, nil) != nil {
		panic("can't start server")
	}
}
