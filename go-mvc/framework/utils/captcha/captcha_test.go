package captcha

import (
	"fmt"
	"testing"
)

// $ go test -v -cover captcha_test.go captcha.go
func TestNewCaptcha(t *testing.T) {
	cp := NewCaptcha(120, 40, 4)
	cp.SetMode(1) // 设置为数学公式
	code, img := cp.OutPut()

	fmt.Println("captcha===", code, img)
}
