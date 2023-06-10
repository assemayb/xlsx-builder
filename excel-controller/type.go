package excelcontroller

type HeadersInfo struct {
	en string
	ar string
}

type RequestBody struct {
	Headers   []HeadersInfo
	Data      [][]interface{}
	lang      string
	sheetName string
}
