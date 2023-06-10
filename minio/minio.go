package minio

import (
	"bytes"
	"fmt"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/tealeg/xlsx"
)

var (
	client *minio.Client
	once   sync.Once
)

func NewClient() (client *minio.Client, err error) {
	endpoint := os.Getenv("minio_endpoint")
	accessKeyID := os.Getenv("minio_access_key")
	secretAccessKey := os.Getenv("minio_secret_key")
	once.Do(func() {
		var err error
		client, err = minio.New(endpoint, &minio.Options{
			Secure:    false,
			Transport: nil,
			Creds:     credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		})
		if err != nil {
			fmt.Println(err)
			return
		}
	})
	return client, nil
}

func PushFileToMiniO(ctx *gin.Context, file *xlsx.File) (string, error) {
	var err error
	var bucketName = "excel"
	var objectName = "myobject"
	var contentType = "application/octet-stream"

	buffer := new(bytes.Buffer)
	err = file.Write(buffer)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	reader := bytes.NewReader(buffer.Bytes())
	// fileSize := int64(buffer.Len())

	fileInfo, err := client.PutObject(ctx, bucketName, objectName, reader, -1, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return fileInfo.Key, nil

}
