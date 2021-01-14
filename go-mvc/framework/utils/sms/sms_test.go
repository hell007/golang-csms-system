package sms

import "testing"

var (
	accessKeyId   = ""
	accessSecret  = ""
	signName      = ""
	templateCode  = ""
	phoneNumber   = ""
	templateParam = map[string]interface{}{"name": "nanjishidu"}
)

// 测试单个文件，一定要带上被测试的原文件，如果原文件有其他引用，也需一并带上
// $ go test -v sms_test.go sms.go //sms.go是依赖包
func TestAlibabaSendSms(t *testing.T) {
	m := NewAlibabaSendRequest(accessKeyId, accessSecret)
	alibabaSendResponse, err := m.AlibabaSendSms(signName, templateCode, phoneNumber, templateParam)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(alibabaSendResponse)
	if alibabaSendResponse.Code != "OK" {
		t.FailNow()
	}
}
