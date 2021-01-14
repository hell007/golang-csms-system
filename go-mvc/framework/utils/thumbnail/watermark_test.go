package thumbnail

import (
	"fmt"
	"go-mvc/framework/conf"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"testing"
)

// 图片添加水印
// $ go test -v watermark_test.go utils.go watermark.go
func TestWaterMark(t *testing.T) {
	fileName := "test.jpg"
	newName := "water-test.jpg"
	newImage, info, err1 := LoadImage(conf.GetUploadFile() + fileName)
	fmt.Println("添加水印1", info, err1)

	m, _ := WaterMark(newImage)

	out, err2 := os.Create(conf.GetUploadFile() + newName)
	if err2 != nil {
		log.Fatal(err2)
	}
	defer out.Close()

	switch info {
	case "jpeg":
		jpeg.Encode(out, m, nil)
	case "png":
		png.Encode(out, m)
	case "gif":
		gif.Encode(out, m, nil)
	}

	fmt.Println("添加水印2", err2)
}
