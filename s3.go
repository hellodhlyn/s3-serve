package main

import (
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var s3Svc *s3.S3

func downloadFromS3(bucketName, key string) (io.ReadCloser, error) {
	out, err := s3Svc.GetObject(&s3.GetObjectInput{Bucket: aws.String(bucketName), Key: aws.String(key)})
	if err != nil {
		return nil, err
	}

	return out.Body, nil
}

func init() {
	sess := session.Must(session.NewSession())
	s3Svc = s3.New(sess, &aws.Config{Region: aws.String(s3BucketRegion)})
}
