package main

import (
	"fmt"
)

func main() {
	m1 := map[string]int{"a": 1, "b": 2}
	fmt.Println(m1)
	fmt.Println(m1["a"])

	// キーの存在確認
	value, exist := m1["a"]
	fmt.Println(value, exist)

	// 要素削除
	delete(m1, "b")
	fmt.Println(m1)

	m2 := make(map[string]int)
	fmt.Println(m2)
	m2["c"] = 3
	fmt.Println(m2)
}
