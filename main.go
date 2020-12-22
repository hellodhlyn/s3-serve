package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/hellodhlyn/s3-serve/sources"
)

var srcs []sources.Source

func getSourceMust(source sources.Source, err error) sources.Source {
	if err != nil {
		panic(err)
	}
	return source
}

func main() {
	srcs = []sources.Source{
		getSourceMust(sources.NewLocalCacheSource(localStoragePath)),
		getSourceMust(sources.NewAWSS3Source(s3BucketName, s3BucketRegion, s3ObjectKeyPrefix)),
	}

	http.HandleFunc("/ping", ping)
	http.HandleFunc("/", serve)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	fmt.Println("Server listing 0.0.0.0:" + port)
	fmt.Println(http.ListenAndServe(":"+port, nil))
}
