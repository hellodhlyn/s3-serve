package main

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

func ping(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("pong"))
}

func serve(w http.ResponseWriter, r *http.Request) {
	objectKey := s3ObjectKeyPrefix + r.URL.Path
	readFromCache := true

	// read file from cache or S3
	body, err := readCache(objectKey)
	if err != nil {
		body, err = downloadFromS3(s3BucketName, objectKey)
		readFromCache = false
	}
	if err != nil {
		log.Printf("Failed to download from s3. (%s)\n%v\n", r.URL.Path, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// write cache if not exists
	if !readFromCache {
		var buf bytes.Buffer
		tee := io.TeeReader(body, &buf)
		body = &buf

		err = writeCache(objectKey, tee)
		if err != nil {
			log.Printf("Failed to store cache.\n%v\n", err)
		}
	}

	w.Header().Set("Cache-Control", "max-age=86400")
	_, err = io.Copy(w, body)
	if err != nil {
		log.Printf("Failed to send response.\n%v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
