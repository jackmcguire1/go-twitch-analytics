package utils

import (
	"github.com/jackmcguire1/go-twitch-analytics/internal/aws/s3"

	"net/http"
	"time"
)

// DownloadFileToS3 write file to S3 from URL
func DownloadFileToS3(bucket, key, url string) (err error) {
	client := http.Client{
		Timeout: time.Minute * 2,
	}

	var resp *http.Response
	resp, err = client.Get(url)
	if err != nil {
		return
	}
	err = s3.WriteFromReader(bucket, key, resp.Body)
	return
}
