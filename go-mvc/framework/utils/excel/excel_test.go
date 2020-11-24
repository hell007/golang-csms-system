package excel

import (
	"testing"
	"go-mvc/framework/conf"
)

func TestImport(t *testing.T) {
	inFile := conf.GetUploadFile() + "test.xlsx"
	Import(inFile)
}

func TestExport(t *testing.T) {
	outFile := conf.GetUploadFile() + "test2.xlsx"
	Export(outFile, "excel测试")
}