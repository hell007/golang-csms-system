package contains

import (
	"fmt"
	"testing"
)

func TestCollectionContain(t *testing.T) {
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
