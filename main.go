package main

import (
	controller "excel-builder/excel-controller"
	"log"

	"github.com/gin-gonic/gin"
)

var (
	server *gin.Engine
)

func init() {
	server = gin.New()
	server.Use(gin.Logger(), gin.Recovery())
}

func main() {
	server.POST("/api/excel/build", controller.CreateExcelFile)
	log.Fatal(server.Run(":9197"))
}
