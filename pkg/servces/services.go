package servces

import (
	"errors"
	"fmt"
	"github.com/burhon94/json/cmd/writer"
	"github.com/google/uuid"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

const MediaUrl = "web/media"

func EnvOrFlag(envName string, flag *string) (value string, ok bool) {
	if flag == nil {
		return *flag, true
	}

	return os.LookupEnv(envName)
}

func SaveFile(file io.Reader, contentType string) (string, error) {
	if len(contentType) <= 0 {
		return "", errors.New("invalid extensions")
	}

	uuidV4 := uuid.New().String()
	fileName := fmt.Sprintf("%s%s", uuidV4, contentType)
	path := filepath.Join(MediaUrl, fileName)

	dstFile, err := os.Create(path)
	if err != nil {
		log.Printf("can't create file: %v", err)
	}
	defer func() {
		if dstFile.Close() != nil {
			log.Print("can't close dstFile")
		}
	}()

	_, err = io.Copy(dstFile, file)
	if err != nil {
		log.Printf("can't save file: %s, error: %v", file, err)
	}

	fPath := strings.Split(path, fileName)
	pathFile := fPath[0]
	uploadFile, err := writer.JsonFileUpload(pathFile)
	if err != nil {
		return "", errors.New("error while convert to JSON")
	}

	return uploadFile, nil
}
