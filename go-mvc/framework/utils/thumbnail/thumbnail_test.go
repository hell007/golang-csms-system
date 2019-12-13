package thumbnail

import (
	"testing"
)

// go test -v -cover thumbnail.go crop.go utils.go thumbnail_test.go
func TestThumbnailCrop(t *testing.T) {
	fileName :=  "water-test.jpg"
	ThumbnailSave(fileName, 400, 400, 7)
}

