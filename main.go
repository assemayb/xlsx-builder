package main

import (
	controller "excel-builder/excel-controller"
	minioPackage "excel-builder/minio"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/minio/minio-go/v7"
)

var (
	server        *gin.Engine
	minioInstance *minio.Client
)

func init() {

	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server = gin.New()
	minioInstance, err = minioPackage.NewClient()

	if err != nil {
		log.Fatal(err)
	}

	server.Use(gin.Logger(), gin.Recovery())
}

func main() {
	server.POST("/api/excel/build", controller.CreateExcelFile)
	log.Fatal(server.Run(":9197"))
}
