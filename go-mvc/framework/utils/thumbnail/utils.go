package thumbnail

import (
	"image"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
	//_ "github.com/golang/image/webp"

	"go-mvc/framework/conf"
)

// 加载图片
func LoadImage(imgPath string) (img image.Image, info string, err error) {
	file, err := os.Open(imgPath)
	if err != nil {
		return
	}
	defer file.Close()
	img, info, err = image.Decode(file)
	return
}

// 拆分参数
func ParseImgArg(imgArg string) (uint, uint) {
	args := strings.Split(imgArg, "x")
	if len(args) != 2 {
		return 0, 0
	}
	width, _ := strconv.Atoi(args[0])
	height, _ := strconv.Atoi(args[1])
	return uint(width), uint(height)
}

//去除名称前缀，获得文件类型,重新命名
func ParseName(fileName string, flag int) string {
	var (
		fileNamePre string
	)
	args := strings.Split(fileName, ".")

	switch flag {
	case 1:
		fileNamePre = conf.GlobalConfig.UploadStyle[0]
	case 2:
		fileNamePre = conf.GlobalConfig.UploadStyle[1]
	case 3:
		fileNamePre = conf.GlobalConfig.UploadStyle[2]
	default: //0
		fileNamePre = ""
	}
	return RenameFile() + fileNamePre + "." + args[1]
}

// 文件重命名
func RenameFile() string {
	stringTime := time.Now().Format("2006-01-02")
	//step1: 设置种子数
	rand.Seed(time.Now().UnixNano())
	//step2：获取随机数
	stringInt64 := strconv.FormatInt(rand.Int63n(1000000000), 10)
	fileName := strings.Replace(stringTime, "-", "", -1)
	return fileName + stringInt64
}
