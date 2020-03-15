package server

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path/filepath"
)

const (
	templatesUrl = "web/templates"
	indexUrl     = "web/templates/index"
	assetsUrl    = "web/assets"
)

func handleRedirect(responseWriter http.ResponseWriter, request *http.Request) {
	http.Redirect(responseWriter, request, "/upload", http.StatusFound)
}

func (receiver *Server) handleUpload() func(http.ResponseWriter, *http.Request) {
	var (
		tpl *template.Template
		err error
	)

	tpl, err = template.ParseFiles(
		filepath.Join(indexUrl, "index.gohtml"),
		filepath.Join(templatesUrl, "base.gohtml"),
	)
	if err != nil {
		panic(err)
	}
	return func(writer http.ResponseWriter, request *http.Request) {

		data := struct {
			Title string
		}{
			Title: "Uploader File",
		}

		err = tpl.Execute(writer, data)
		if err != nil {
			log.Print(err)
			http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

	}
}

func (receiver *Server) handleFavicon() func(http.ResponseWriter, *http.Request) {
	file, err := ioutil.ReadFile(filepath.Join(assetsUrl, "favicon.ico"))
	if err != nil {
		panic(err)
	}
	return func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write(file)
		if err != nil {
			log.Print(err)
		}
	}
}
