package main

import (
	"fmt"
)

// golang に while は無い
func main() {
	// 普通の for
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		} else if i == 8 {
			break
		}
		fmt.Print(i)
		fmt.Print(", ")
	}
	fmt.Println()

	// while 的な
	n := 0
	for n < 10 {
		fmt.Print(n)
		fmt.Print(", ")
		n++
	}
	fmt.Println()

	// range
	s := []string{"a", "b", "c"}
	for i, v := range s {
		fmt.Print(i)
		fmt.Print(", ")
		fmt.Println(v)
	}

	// ブランク修飾子
	// _にしておくと何が入ってきてもそれを廃棄してくれるという仕様
	for _, v := range s {
		fmt.Println(v)
	}

	// マップ
	m := map[string]int{"x": 10, "y": 20}
	for k, v := range m {
		fmt.Println(k, v)
	}
}
