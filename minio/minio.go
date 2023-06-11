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

// ?? fix this
func NewClient() (client *minio.Client, err error) {
	endpoint := os.Getenv("minio_endpoint")
	accessKeyID := os.Getenv("minio_access_key")
	secretAccessKey := os.Getenv("minio_secret_key")
	once.Do(func() {
		var error error
		client, error = minio.New(endpoint, &minio.Options{
			Secure:    false,
			Transport: nil,
			Creds:     credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		})

		if error != nil {
			fmt.Println(error)
		}
	})
	return client, nil
}

func PushFileToMiniO(ctx *gin.Context, file *xlsx.File) (string, error) {

	fmt.Println("  ------->>  Pushing file to minio  <<-------  ")
	// context := context.Background()
	var err error
	var bucketName = "excel"
	var objectName = "myObject"
	var contentType = "application/octet-stream"

	var fileBuffer = new(bytes.Buffer)
	err = file.Write(fileBuffer)
	if err != nil {
		fmt.Println("Error Writing file data to a buffer", err)
		panic(err)

	}

	reader := bytes.NewReader(fileBuffer.Bytes())
	fileSize := int64(fileBuffer.Len())

	clientInstance, err := NewClient()
	if err != nil {
		fmt.Println("Error creating minio client")
		fmt.Println(err)
	}
	fmt.Println(clientInstance.IsOnline())
	fileInfo, err := clientInstance.PutObject(
		ctx,
		bucketName,
		objectName,
		reader,
		fileSize,
		minio.PutObjectOptions{ContentType: contentType},
	)

	if err != nil {
		fmt.Println("Error uploading file to minio")
		fmt.Println(err)
	}
	return fileInfo.Key, nil
}
