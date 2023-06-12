package minio

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/tealeg/xlsx"
)

var (
	lock           sync.Mutex
	clientInstance *minio.Client
	once           sync.Once
)

func SetMinioClientConnection(endpoint, accessKeyID, secretAccessKey string) (client *minio.Client, err error) {
	once.Do(func() {
		var error error
		clientInstance, error = minio.New(endpoint, &minio.Options{
			Secure:    false,
			Transport: nil,
			Creds:     credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		})
		if error != nil {
			fmt.Println("Error creating minio client")
			panic(error)
		}
	})
	return client, nil
}

func GetInstance() *minio.Client {
	if clientInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if clientInstance == nil {
			var error error
			endpoint := os.Getenv("minio_endpoint")
			accessKeyID := os.Getenv("minio_access_key")
			secretAccessKey := os.Getenv("minio_secret_key")
			clientInstance, error = SetMinioClientConnection(endpoint, accessKeyID, secretAccessKey)
			if error != nil {
				fmt.Println("Error creating minio client")
				panic(error)
			}
		}
	}
	return clientInstance
}

func PushFileToMiniO(ctx *gin.Context, file *xlsx.File) (minio.UploadInfo, error) {
	context := context.Background()
	bucketName := "excel"
	objectName := "myObject.xlsx"
	contentType := "application/octet-stream"

	var err error
	var fileBuffer = new(bytes.Buffer)
	err = file.Write(fileBuffer)
	if err != nil {
		fmt.Println("Error Writing file data to a buffer", err)
		panic(err)
	}
	reader := bytes.NewReader(fileBuffer.Bytes())
	fileSize := int64(fileBuffer.Len())

	fileInfo, err := GetInstance().PutObject(
		context,
		bucketName,
		objectName,
		reader,
		fileSize,
		minio.PutObjectOptions{ContentType: contentType},
	)
	if err != nil {
		fmt.Println(err)
	}
	return fileInfo, nil
}
