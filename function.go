package main

import "fmt"

/*
- 引数の型、返り値の型は必須
- 複数の返り値を返す事ができる
- 関数もデータ型なので、変数に格納可能
- 即時関数の様に宣言して、その場で実行する事も可能
*/

func sayName(name string) {
	fmt.Println(name)
}

func getMessage(name string) string {
	msg := "hi! my name is " + name
	return msg
}

func getHelloMessage(name string) (msg string) {
	msg = "Hello " + name
	return
}

func main() {
	sayName("pepese")
	fmt.Println(getMessage("pepese"))
	fmt.Println(getHelloMessage("pepese"))

	f := func(name string) string {
		return getMessage(name)
	}
	fmt.Println(f("pepese"))

	func(name string) {
		fmt.Println(name)
	}("pepese")
}
