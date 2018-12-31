package main

import (
	"fmt"
)

/*
defer 文を使うことで、それを呼び出した関数が終了する際に実行すべき処理を記述することができる
*/
func func1() {
	defer fmt.Println("func1 End")
	fmt.Println("func1 Stert")
}

/*
panic() した時にも呼び出されるので recover() する
*/
func func2() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("func2 Recover!:", err)
		}
		fmt.Println("func2 End")
	}()

	fmt.Println("func2 Start")
	panic("func2 Panic!")
}

/*
defer 文は return の後に実行される
*/
func func3() (int, error) {
	defer fmt.Println("func3 End")
	fmt.Println("func3 Stert")
	return fmt.Println("func3 string")
}

func main() {
	func1()
	func2()
	//fmt.Println(func3())
	_, _ = func3()
}
