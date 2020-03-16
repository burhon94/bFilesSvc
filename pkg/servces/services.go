package servces

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const MediaUrl = "web/media"

func EnvOrFlag(envName string, flag *string) (value string, ok bool) {
	if flag == nil {
		return *flag, true
	}

	return os.LookupEnv(envName)
}

func SaveFile(w http.ResponseWriter, file io.Reader, contentType string) (string, error) {
	if len(contentType) <= 0 {
		return "", errors.New("invalid extensions")
	}

	uuidV4 := uuid.New().String()
	fileName := fmt.Sprintf("%s%s", uuidV4, contentType)
	path := filepath.Join(MediaUrl, fileName)
	_, err := os.Stat(MediaUrl)
	if os.IsNotExist(err) {
		err := os.Mkdir(MediaUrl, 0777)
		if err != nil {
			panic("can't create dir")
		}
	}

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

	return fileName, nil
}
