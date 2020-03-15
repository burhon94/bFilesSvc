package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func (receiver *Server) GorillaInit(addr string) {
	router := mux.NewRouter()
	router.HandleFunc("/", handleRedirect)
	router.HandleFunc("/upload", receiver.handleUploadPage())
	router.HandleFunc("/favicon.ico", receiver.handleFavicon())

	//get files from media dir
	router.HandleFunc("/media", http.StripPrefix("/media", http.FileServer(http.Dir(MediaUrl))).ServeHTTP)

	http.Handle("/", router)
	fmt.Println("Server is listening...")
	http.ListenAndServe(addr, nil)
}
