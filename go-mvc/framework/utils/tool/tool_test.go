package tool

import (
	"fmt"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	//c := 1
	//c := ""
	c := []int{} //slice
	//c := []int{1, 2, 3} //slice

	res := IsEmpty(c)
	fmt.Println("res==", res)
}

func TestContain(t *testing.T) {
	//数组
	a := 1
	b := [3]int{1, 2, 3}
	fmt.Println(Contain(b, a))

	//slice
	c := 1
	d := []int{1, 2, 3}
	fmt.Println(Contain(d, c))

	//map
	var e = map[int]string{1: "1", 2: "2"}
	fmt.Println(Contain(e, 2))
}
