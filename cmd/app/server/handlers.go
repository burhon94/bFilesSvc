package server

import (
	"encoding/json"
	"github.com/burhon94/bFilesSvc/pkg/resposes"
	"github.com/burhon94/bFilesSvc/pkg/servces"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"path/filepath"
	"strings"
)

const (
	MediaUrl = "web/media"

	multipartMaxBytes = 10 * 1024 * 1024
)

type fileStruct struct {
	FileName string `json:"fileName"`
}


func (receiver *Server) handlerRespHealth() func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		_, err := writer.Write([]byte("ok"))
		if err != nil {
			resposes.InternalServerError(writer, err)
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
		filesList := make([]fileStruct, 0)

		for _, file := range files["files"] {
			contentType := path.Ext(file.Filename)
			openFile, err := file.Open()
			if err != nil {
				log.Printf("can't create file: %v", err)
				continue
			}

			uploadedFiles, err = servces.SaveFile(writer, openFile, contentType)
			if err != nil {
				log.Printf("can't save file: %v", err)
				continue
			}
			filesList = append(filesList, fileStruct{
				FileName: uploadedFiles,
			})
		}

		bytes, err := json.Marshal(filesList)
		writer.Header().Set("Content-Type", "application/json")
		_, err = writer.Write(bytes)
		if err != nil {
			resposes.InternalServerError(writer, err)
		}

	}
}

func (receiver *Server) handleGetFile() func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		fileName := strings.TrimPrefix(request.RequestURI, "/api/files/")

		file, err := ioutil.ReadFile(filepath.Join(MediaUrl, fileName))
		if err != nil {
			resposes.BadRequest(err, writer)
			return
		}

		_, err = writer.Write(file)
		if err != nil {
			resposes.InternalServerError(writer, err)
			return
		}
	}
}
