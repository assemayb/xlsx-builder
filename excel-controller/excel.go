package excelcontroller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tealeg/xlsx"
)

func CreateExcelFile(ctx *gin.Context) {
	var body RequestBody
	err := ctx.BindJSON(&body)
	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	file := buildExcel(body.Data, body.Headers, body.Lang, body.SheetName)
	err = file.Save(fmt.Sprintf("%s.xlsx", body.SheetName))

	if err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "success"})
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
			cell.Value = header.En
		} else {
			cell.Value = header.Ar
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
