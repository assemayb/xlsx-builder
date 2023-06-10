package minio

import (
	"fmt"
	"sync"

	"github.com/minio/minio-go/v7"
	"github.com/tealeg/xlsx"
)

var (
	client *minio.Client
	once   sync.Once
)

func NewClient(config MinioConfig) (*minio.Client, error) {
	once.Do(func() {
		var err error
		client, err = minio.New(config.Endpoint, &minio.Options{
			Secure:    false,
			Region:    "us-east-1",
			Transport: nil,
		})
		if err != nil {
			fmt.Println("Error initializing MinIO client:", err)
		}
	})
	return client, nil
}

func PushFileToMiniO(file *xlsx.File) (string, error) {
	return "", nil
}
