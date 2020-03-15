package resposes

import (
	"log"
	"net/http"
)

func BadRequest(err error, w http.ResponseWriter)  {
	log.Print(err)
	http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
}

func InternalServerError(w http.ResponseWriter, err error) {
	log.Print(err)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	return
}
