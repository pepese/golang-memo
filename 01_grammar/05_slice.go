package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	fmt.Println(s1)      // [1 2 3]
	fmt.Println(s1[0:1]) // [1]
	fmt.Println(s1[1:])  // [2 3]
	fmt.Println(len(s1)) // [3]

	// 要素の追加
	s1 = append(s1, 4, 5)
	fmt.Println(s1)      // [1 2 3 4 5]
	fmt.Println(len(s1)) // 5

	// make
	// データ型のデフォルト値で指定サイズで作成
	s2 := make([]int, len(s1))
	fmt.Println(s2) // [0 0 0 0 0]
	// コピー
	s3 := copy(s2, s1) // 返り値は要素数
	fmt.Println(s2)    // [1 2 3 4 5]
	fmt.Println(s3)    // 5
}
