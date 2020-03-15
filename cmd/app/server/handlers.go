package server

import (
	"fmt"
	"net/http"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Nice Job")
	w.Header().Set("Status Code", "200")
}
