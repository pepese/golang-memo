package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	fmt.Println(s1)
	fmt.Println(s1[0:1])
	fmt.Println(s1[1:])
	fmt.Println(len(s1))

	// 要素の追加
	s1 = append(s1, 4, 5)
	fmt.Println(s1)
	fmt.Println(len(s1))

	// make
	// データ型のデフォルト値で指定サイズで作成
	s2 := make([]int, len(s1))
	fmt.Println(s2)
	// コピー
	s3 := copy(s2, s1)
	fmt.Println(s2)
	fmt.Println(s3)
}
