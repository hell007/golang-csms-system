package thumbnail

import (
	"encoding/base64"
	"go-mvc/framework/utils/files"
	"os"
	"strings"
)

// base64图片上传
func Base64Upload(fileDir string, strImage string) (string, error) {
	var (
		err      error
		image    string
		fileType string
		fireName string
		filePath string
	)

	image, fileType = splitPictureCodeAndType(strImage)
	fireName = RenameFile() + "." + fileType
	filePath, err = files.MakeFilePath(fileDir, fireName)
	if err != nil {
		return "", nil
	}

	err = base642Image(filePath, image)
	if err != nil {
		return "", nil
	}
	return filePath, err
}

// 切割 base64 image sting and image type
func splitPictureCodeAndType(data string) (string, string) {
	var (
		a         []string
		b         []string
		c         []string
		image     string
		imageType string
	)

	//data := "data:image/jpg;base64,adsfsdfsdf"
	a = strings.Split(data, ",")
	image = a[1]
	b = strings.Split(a[0], ";")
	c = strings.Split(b[0], "/")
	imageType = c[1]
	return image, imageType
}

// base64转为图片
func base642Image(fileName string, strImage string) error {
	decode, err := base64.StdEncoding.DecodeString(strImage)
	file, err := os.Create(fileName)
	defer file.Close()

	_, err = file.Write(decode)
	return err
}
