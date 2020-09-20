package main

import (
	"io"
	"log"
	"net/http"
)

func ping(w http.ResponseWriter, _ *http.Request) {
	_, _ = w.Write([]byte("pong"))
}

func serve(w http.ResponseWriter, r *http.Request) {
	objectKey := s3ObjectKeyPrefix + r.URL.Path
	body, err := downloadFromS3(s3BucketName, objectKey)
	if err != nil {
		log.Printf("Failed to download from s3. (%s)\n%v\n", r.URL.Path, err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, body)
	if err != nil {
		log.Printf("Failed to send response.\n%v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}
