package main

import (
	"errors"
	"os"
)

var (
	s3BucketName      string
	s3BucketRegion    string
	s3ObjectKeyPrefix string

	localStoragePath string
)

func getEnvMust(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	panic(errors.New("environment variable require: " + key))
}

func init() {
	// aws_s3
	s3BucketName = getEnvMust("S3_BUCKET_NAME")
	s3BucketRegion = getEnvMust("S3_BUCKET_REGION")
	s3ObjectKeyPrefix = getEnvMust("S3_OBJECT_KEY_PREFIX")

	// local_storage
	localStoragePath = os.Getenv("LOCAL_STORAGE_PATH")
	if _, err := os.Stat(localStoragePath); localStoragePath != "" && os.IsNotExist(err) {
		_ = os.MkdirAll(localStoragePath, 0755)
	}
}
