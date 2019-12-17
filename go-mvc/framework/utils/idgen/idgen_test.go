package idgen

import (
	"fmt"
	"testing"
)

//func TestGenerate(t *testing.T) {
//	Generate()
//}

//func TestGenerateUuid(t *testing.T) {
//	s := GenerateUuid()
//	fmt.Println("TestGenerateUuid====", s)
//}

func TestGenerateOrdersn(t *testing.T) {
	osn := GenerateOrdersn()
	fmt.Println("TestGenerateOrdersn====", osn)
}
