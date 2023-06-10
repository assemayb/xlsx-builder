package excelcontroller

type HeadersInfo struct {
	En string `json:"en"`
	Ar string `json:"ar"`
}

type RequestBody struct {
	Headers   []HeadersInfo   `json:"headers"`
	Data      [][]interface{} `json:"data"`
	Lang      string          `json:"lang"`
	SheetName string          `json:"sheetName"`
}
