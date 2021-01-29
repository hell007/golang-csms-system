package encrypt

import (
	"fmt"
	"testing"
)

const pwd = "admin123"

//const keys = "0123456789012345" //16
//const keys = "TESTWdczx13213EDDSWQ!!@W" //24
//const keys = "TESTECRETARRSWdczx13213EDDSWQ!@W" //32
const keys = "TESTWdczx13213EDDSWQ!!@W"

func TestAESDecrypt(t *testing.T) {
	password := AESEncrypt(pwd, keys)
	fmt.Println("加密 ==：", password)
}

func TestAESEncrypt(t *testing.T) {
	password := AESEncrypt(pwd, keys)
	pwd2 := AESDecrypt(password, keys)
	fmt.Println("解密 ==：", pwd2)
}

func TestCheckPWD(t *testing.T) {
	password := AESEncrypt(pwd, keys)
	b := CheckPWD(pwd, password, keys)
	fmt.Println("相等？ ==：", b)
}
