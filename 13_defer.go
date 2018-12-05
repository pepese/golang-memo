package main

import (
	"fmt"
)

/*
defer 文を使うことで、それを呼び出した関数が終了する際に実行すべき処理を記述することができる
*/
func func1() {
	defer fmt.Println("End")
	fmt.Println("Stert")
}

/*
panic() した時にも呼び出されるので recover() する
*/
func func2() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Recover!:", err)
		}
		fmt.Println("End")
	}()

	fmt.Println("Start")
	panic("Panic!")
}

func main() {
	func1()
	func2()
}
