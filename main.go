package main

import (
	controller "excel-builder/excel-controller"
	minioPackage "excel-builder/minio"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine
)

func init() {
	var endpoint = os.Getenv("minio_endpoint")
	var accessKeyID = os.Getenv("minio_access_key")
	var secretAccessKey = os.Getenv("minio_secret_key")
	server = gin.New()
	_, err := minioPackage.SetMinioClientConnection(endpoint, accessKeyID, secretAccessKey)
	if err != nil {
		log.Fatal("Minio Error", err)
	}
	server.Use(gin.Logger(), gin.Recovery())
}

func main() {
	server.POST("/api/excel/build", controller.CreateExcelFile)
	log.Fatal(server.Run(":9197"))
}
