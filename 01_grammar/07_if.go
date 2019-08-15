package main

import "fmt"

func main() {
	// 普通の条件分岐
	score := 1
	if score > 80 {
		fmt.Println("1")
	} else if score > 60 {
		fmt.Println("2")
	} else {
		fmt.Println("3") // !
	}

	// if 文内でしか使わない変数は if 文の中で定義できる
	if a := 12; a > 80 {
		fmt.Println("1")
	} else if a > 60 {
		fmt.Println("2")
	} else {
		fmt.Println("3") // !
	}
}
