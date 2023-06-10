package minio

import "github.com/minio/minio-go/v7"

type MinioConfig struct {
	Endpoint        string
	AccessKeyID     string
	SecretAccessKey string
	UseSSL          bool
}

type MiniClientType *minio.Client
