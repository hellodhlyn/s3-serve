package main

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/hellodhlyn/s3-serve/sources"
)

func ping(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("pong"))
}

func readFromSources(key string, targetSrcs []sources.Source) (io.Reader, error) {
	currentSource := targetSrcs[0]
	body, err := currentSource.Read(key)
	if body != nil && err == nil {
		return body, nil
	}

	if len(targetSrcs) == 1 {
		return nil, errors.New("no such file")
	}

	body, err = readFromSources(key, targetSrcs[1:])
	if err == nil {
		var buf bytes.Buffer
		tee := io.TeeReader(body, &buf)
		body = &buf
		err = currentSource.Write(key, tee)
		if err != nil {
			fmt.Printf("failed to write file: %s\n%v\n", key, err)
		}
	}

	return body, err
}

func serve(w http.ResponseWriter, r *http.Request) {
	body, err := readFromSources(r.URL.Path, srcs)
	if err != nil {
		log.Printf("failed to read file: %s\n%v\n", r.URL.Path, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Cache-Control", "max-age=86400")
	_, err = io.Copy(w, body)
	if err != nil {
		log.Printf("Failed to send response.\n%v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
