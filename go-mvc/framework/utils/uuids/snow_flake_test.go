package uuids

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestSnowFlake(t *testing.T) {
	// 生成节点实例
	worker, err := NewWorker(1)

	if err != nil {
		fmt.Println(err)
		return
	}

	ch := make(chan int64)
	count := 10000
	// 并发 count 个 goroutine 进行 snowflake ID 生成
	for i := 0; i < count; i++ {
		go func() {
			id := worker.GetId()
			fmt.Println("i==>", id)
			ch <- id
		}()
	}

	defer close(ch)

	m := make(map[int64]int)
	for i := 0; i < count; i++  {
		id := <- ch
		// 如果 map 中存在为 id 的 key, 说明生成的 snowflake ID 有重复
		_, ok := m[id]
		if ok {
			t.Error("ID is not unique!\n")
			return
		}
		// 将 id 作为 key 存入 map
		m[id] = i
	}
	// 成功生成 snowflake ID
	fmt.Println("All", count, "snowflake ID Get successed!")
}

func TestGetId(t *testing.T) {
	// 生成节点实例
	worker, _ := NewWorker(1)
	id := worker.GetId()
	fmt.Println("id==", id, len(strconv.FormatInt(id, 10)))


	t2 := time.Now().Unix()
	fmt.Println("now==>", t2)

	var timeUnix int64 = 1525705533
	fmt.Println("1525705533==>", time.Unix(timeUnix,0))

	t3 := time.Date(2021, 1, 1, 1, 0, 0, 0, time.Local).Unix()
	fmt.Println("t3==>", t3)
}