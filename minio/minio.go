package minio

import (
	"fmt"
	"os"
	"sync"

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
	fmt.Println("=====================================")
	fmt.Println("=====================================")
	fmt.Println("=====================================")
	fmt.Println(endpoint, accessKeyID, secretAccessKey)
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

func PushFileToMiniO(file *xlsx.File) (string, error) {
	return "", nil
}
