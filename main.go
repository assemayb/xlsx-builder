package main

import (
	controller "excel-builder/excel-controller"
	minioPackage "excel-builder/minio"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var (
	server *gin.Engine
)

func init() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	server = gin.New()
	_, err = minioPackage.SetMinioClientConnection()
	if err != nil {
		log.Fatal("Minio Error", err)
	}
	server.Use(gin.Logger(), gin.Recovery())
}

func main() {
	server.POST("/api/excel/build", controller.CreateExcelFile)
	log.Fatal(server.Run(":9197"))
}
