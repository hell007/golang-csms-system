package encrypt

import (
	"fmt"
	"testing"
)

const pwd = "123456"

//const keys = "0123456789012345" //16
//const keys = "ARRSWdczx13213EDDSWQ!!@W" //24
//const keys = "JIESECRETARRSWdczx13213EDDSWQ!@W" //32
const keys = "JIE_SECRET_KEY2019-10-21"

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
