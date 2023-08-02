package controllers

import (
	"fmt"
	"io"
	"mime/multipart"
	"sima/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

var bucketName string = "sima-bucket"

func uploadFile(fileHeader *multipart.FileHeader) (string, error) {
	file, err := fileHeader.Open()
	if err != nil {
		return "", err
	}
	defer file.Close()

	UploadedURL, err := saveFile(file, fileHeader)
	if err != nil {
		return "", err
	}

	return UploadedURL, nil
}

func saveFile(fileReader io.Reader, fileHeader *multipart.FileHeader) (string, error) {

	_, err := config.Uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileHeader.Filename),
		Body:   fileReader,
	})
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucketName, fileHeader.Filename)

	return url, nil
}
