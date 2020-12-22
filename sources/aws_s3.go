package sources

import (
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type AWSS3Source struct {
	bucketName      string
	objectKeyPrefix string

	s3Client *s3.S3
}

func (src *AWSS3Source) Read(key string) (io.Reader, error) {
	out, err := src.s3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(src.bucketName),
		Key:    aws.String(src.objectKeyPrefix + key),
	})
	if err != nil {
		return nil, err
	}

	return out.Body, nil
}

func (src *AWSS3Source) Write(_ string, _ io.Reader) error {
	// not implmented
	return nil
}

func NewAWSS3Source(bucketName, bucketRegion, objectKeyPrefix string) (*AWSS3Source, error) {
	sess := session.Must(session.NewSession())
	s3Client := s3.New(sess, &aws.Config{Region: aws.String(bucketRegion)})

	return &AWSS3Source{
		bucketName:      bucketName,
		objectKeyPrefix: objectKeyPrefix,
		s3Client:        s3Client,
	}, nil
}
