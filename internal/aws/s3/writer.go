package s3

import (
	"bytes"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	internal "github.com/jackmcguire1/go-twitch-analytics/internal/aws"
)

var uploader = s3manager.NewUploader(internal.Session)

// Write inserts file from bytes onto S3
func Write(bucket, key string, data []byte) (err error) {
	_, err = uploader.Upload(&s3manager.UploadInput{
		Body:   bytes.NewReader(data),
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	return
}

//WriteFromReader inserts file onto S3 from reader
func WriteFromReader(bucket, key string, r io.Reader) (err error) {
	_, err = uploader.Upload(&s3manager.UploadInput{
		Body:   r,
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	return
}
