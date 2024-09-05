package storage

import (
	"algvisual/internal/infrastructure/config"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Storage struct {
	c config.AppConfig
}

func (s S3Storage) CreateSession() (*session.Session, error) {
	region := s.c.S3Config.Region

	accessKey := s.c.S3Config.AccessKeyID
	secretKey := s.c.S3Config.SecretKeyID

	return session.NewSession(&aws.Config{
		Region: aws.String(region),
		Credentials: credentials.NewStaticCredentials(
			accessKey,
			secretKey,
			"",
		),
	})
}

func (s S3Storage) Upload() (string, error) {
	return "", nil
}

func (s S3Storage) UploadFileToS3(
	key, bucket, prefix string,
	file io.ReadSeeker,
	size int64,
) error {
	sess, err := s.CreateSession()
	if err != nil {
		return err
	}
	svc := s3.New(sess)
	objectKey := fmt.Sprintf("%s/%s", prefix, key)
	_, err = svc.PutObject(&s3.PutObjectInput{
		Bucket:        aws.String(bucket),
		Key:           aws.String(objectKey),
		Body:          file,
		ContentLength: aws.Int64(size),
	})
	if err != nil {
		return err
	}

	return nil
}
