package files

import (
	"fmt"
	"runtime"
	"testing"
)

var (
	UploadFile string
)

func GetUploadFile() string {
	sysType := runtime.GOOS
	if sysType == "darwin" {
		UploadFile = "/Users/wzh/Development/git/go/golang-csms-system/golang-mvc/uploads/"
	} else if sysType == "windows" {
		UploadFile = "D:/Dev/cygwin/work/golang/golang-csms-system/golang-mvc/uploads/"
	}
	return UploadFile
}

// 判断文件/或者文件目录 存在否
// $ go test -v -cover files_test.go files.go
func TestFileExists(t *testing.T) {
	//file := "test.jpg"
	//b := FileExists(conf.GetUploadFile() + file)
	//fmt.Println("b===", b)
}

// 创建文件目录
// $ go test -v -cover files_test.go files.go
func TestEnsureDir(t *testing.T) {
	//dir := "test2/"
	//err := EnsureDir(conf.GetUploadFile() + dir)
	//fmt.Println("err===", err)
}

// CreateFile
// $ go test -v -cover files_test.go files.go
func TestCreateFile(t *testing.T) {
	file, err := CreateFile(GetUploadFile() + "readme.txt")
	fmt.Println("===", file, err)
}

// 删除文件
// $ go test -v -cover files_test.go files.go
func TestRemoveFile(t *testing.T) {
	//file := "test.jpg"
	//RemoveFile(GetUploadFile() + file)
}
