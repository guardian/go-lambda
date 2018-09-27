package aws

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func CreateLambda() {
}

func UploadFiles(
	client *s3.S3,
	bucket string,
	pathsToKeys map[string]string,
) error {
	uploader := s3manager.NewUploaderWithClient(client)

	for path, key := range pathsToKeys {
		f, err := os.Open(path)
		if err != nil {
			return err
		}

		_, err = uploader.Upload(&s3manager.UploadInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(key),
			Body:   f,
			ACL:    aws.String("bucket-owner-full-control"),
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func GetClient() (*s3.S3, error) {
	creds := credentials.NewChainCredentials(
		[]credentials.Provider{
			&credentials.SharedCredentialsProvider{
				Filename: "",
				Profile:  "deployTools",
			},
			&credentials.EnvProvider{},
		})

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("eu-west-1"),
		Credentials: creds,
	})
	if err != nil {
		return nil, err
	}

	_, err = sess.Config.Credentials.Get()
	if err != nil {
		return nil, err
	}

	return s3.New(sess), nil
}
