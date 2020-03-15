package server

import (
	"github.com/burhon94/bFilesSvc/pkg/resposes"
	"github.com/burhon94/bFilesSvc/pkg/servces"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"path/filepath"
)

const (
	templatesUrl = "web/templates"
	indexUrl     = "web/templates/index"
	assetsUrl    = "web/assets"
	MediaUrl     = "web/media"

	multipartMaxBytes = 10 * 1024 * 1024
)

func handleRedirect(responseWriter http.ResponseWriter, request *http.Request) {
	http.Redirect(responseWriter, request, "/upload", http.StatusFound)
}

func (receiver *Server) handleUploadPage() func(http.ResponseWriter, *http.Request) {
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

		err = tpl.Execute(writer, struct { }{})
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

func (receiver *Server) handleUploading() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {

		err := request.ParseMultipartForm(multipartMaxBytes)
		if err != nil {
			resposes.BadRequest(err, writer)
			return
		}

		uploadedFiles := ""
		formFiles := request.MultipartForm
		files := formFiles.File

		for _, file := range files["files"] {
			contentType := path.Ext(file.Filename)
			openFile, err := file.Open()
			if err != nil {
				log.Printf("can't create file: %v", err)
				continue
			}

			uploadedFiles, err = servces.SaveFile(openFile, contentType)
			if err != nil {
				log.Printf("can't save file: %v", err)
				continue
			}
		}

		writer.Header().Set("Content-Type", "application/json")
		_, err = writer.Write([]byte(uploadedFiles))
		if err != nil {
			resposes.InternalServerError(writer, err)
		}

		http.Redirect(writer, request, "/upload", http.StatusFound)
	}
}
