package main

import (
	"errors"
	"os"
)

var (
	s3BucketName      string
	s3BucketRegion    string
	s3ObjectKeyPrefix string

	cachePath string
)

func getEnvMust(key string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	panic(errors.New("environment variable require: " + key))
}

func init() {
	s3BucketName = getEnvMust("S3_BUCKET_NAME")
	s3BucketRegion = getEnvMust("S3_BUCKET_REGION")
	s3ObjectKeyPrefix = getEnvMust("S3_OBJECT_KEY_PREFIX")

	home, _ := os.UserHomeDir()
	cachePath = home + "/.cache/s3_serve"
	if _, err := os.Stat(cachePath); os.IsNotExist(err) {
		_ = os.MkdirAll(cachePath, 0755)
	}
}
