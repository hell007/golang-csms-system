package thumbnail

// https://github.com/nfnt/resize.git
// 参数中最大长宽修改为最小长宽
import (
	"errors"
	"github.com/nfnt/resize"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"

	"go-mvc/framework/conf"
)

// 缩略图按照指定的宽和高非失真缩放裁剪
func ThumbnailCrop(minWidth, minHeight uint, img image.Image) image.Image {
	origBounds := img.Bounds()
	origWidth := uint(origBounds.Dx())
	origHeight := uint(origBounds.Dy())
	newWidth, newHeight := origWidth, origHeight

	// Return original image if it have same or smaller size as constraints
	if minWidth >= origWidth && minHeight >= origHeight {
		return img
	}

	if minWidth > origWidth {
		minWidth = origWidth
	}

	if minHeight > origHeight {
		minHeight = origHeight
	}

	// Preserve aspect ratio
	if origWidth > minWidth {
		newHeight = uint(origHeight * minWidth / origWidth)
		if newHeight < 1 {
			newHeight = 1
		}
		newWidth = minWidth
	}

	if newHeight < minHeight {
		newWidth = uint(newWidth * minHeight / newHeight)
		if newWidth < 1 {
			newWidth = 1
		}
		newHeight = minHeight
	}

	thumbImg := resize.Resize(newWidth, newHeight, img, resize.Lanczos3)
	return CropImg(thumbImg, int(minWidth), int(minHeight))
}

// 简单的缩放,指定最大宽和高
func ThumbnailSimple(maxWidth, maxHeight uint, img image.Image) image.Image {
	oriBounds := img.Bounds()
	oriWidth := uint(oriBounds.Dx())
	oriHeight := uint(oriBounds.Dy())

	if maxWidth == 0 {
		maxWidth = oriWidth
	}

	if maxHeight == 0 {
		maxHeight = oriHeight
	}
	return resize.Thumbnail(maxWidth, maxHeight, img, resize.Lanczos3)
}

// 缩略图保存
func ThumbnailSave(filename string, width uint, height uint, flag int) string {
	var (
		filePath string
		newName  string
	)

	switch flag {
	case 1, 2, 3: //  uploads/gooods/
		filePath = conf.GetUploadFile() + conf.GlobalConfig.UploadPicPath[0]
	case 4: // uploads/images/goods/
		filePath = conf.GetUploadFile() + conf.GlobalConfig.UploadEditor[0]
	case 5: // uploads/artilce/
		filePath = conf.GetUploadFile() + conf.GlobalConfig.UploadPicPath[1]
	case 6: // uploads/images/artilce/
		filePath = conf.GetUploadFile() + conf.GlobalConfig.UploadEditor[1]
	case 7: // uploads/ads/
		filePath = conf.GetUploadFile() + conf.GlobalConfig.UploadPicPath[2]
	default:
		filePath = ""
	}

	img, info, err1 := LoadImage(filePath + filename)
	if err1 != nil {
		log.Fatal(err1)
	}

	m := ThumbnailSimple(width, height, img)
	newName = ParseName(filename, flag)

	out, err := os.Create(filePath + newName)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	switch info {
	case "jpeg":
		_ = jpeg.Encode(out, m, nil)
	case "png":
		_ = png.Encode(out, m)
	case "gif":
		_ = gif.Encode(out, m, nil)
	default:
		_ = errors.New("ERROR FORMAT")
	}
	return newName
}
