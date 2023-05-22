package main

import "fmt"

func main() {
	// doMain()
	testMap()
}

func testMap() {
	m := make(map[int]string)
	m[1] = "foo"
	m[2] = "bar"

	for _, v := range m {
		// 以下代码会导致编译错误
		addr := &v
		fmt.Println(v)
		fmt.Println(addr)
	}
}
