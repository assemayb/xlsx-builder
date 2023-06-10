package excelcontroller

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
)

type HeadersInfo struct {
	en string
	ar string
}

func CreateExcelFile(ctx *gin.Context) {
	log.Println("GenerateExcelData")
	var body map[string]interface{}
	err := ctx.BindJSON(&body)
	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("body", body)
	ctx.JSON(200, gin.H{"message": "success"})
	// buildExcelFile([][]interface{}{}, []HeadersInfo{}, "en", "sheetName")
}

func buildExcel(data [][]interface{}, headers []HeadersInfo, lang string, sheetName string) *xlsx.File {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet(sheetName)
	if err != nil {
		log.Fatal(err)
	}

	headerRow := sheet.AddRow()
	for _, header := range headers {
		cell := headerRow.AddCell()
		if lang == "en" {
			cell.Value = header.en
		} else {
			cell.Value = header.ar
		}
	}

	for _, row := range data {
		newRow := sheet.AddRow()
		newRow.Sheet.SetColWidth(0, len(headers), 25)
		for _, cellValue := range row {
			cell := newRow.AddCell()
			cell.Value = cellValue.(string)
		}
	}
	return file
}
