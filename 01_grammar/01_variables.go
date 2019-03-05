package main

import (
	"fmt"
	"reflect"
)

func main() {
	/*
	 * 変数
	 */
	// 初期値無し
	var msg string
	msg = "Hello World"

	// 初期値無し、複数同じ型
	var a, b int
	a, b = 10, 15

	// 初期値あり
	var msg2 string = "Good Night"

	// 初期値有り、値から型を予測できる場合
	var msg3 = "Good Morning"

	// 初期値有り、値から型を予測できる場合、var 省略
	msg4 := "Good Evening"

	// まとめて定義（ := は使えない。 var あるから）
	var (
		c    int                // 初期値無し
		d, e float64            // 初期値無し、複数同じ型
		f    int     = 1        // 初期値有り
		g            = "fuge"   // 初期値有り、型省略
		h, i         = 0.0, 1.1 // 初期値有り、複数
	)
	fmt.Println(msg, msg2, msg3, msg4, a, b, c, d, e, f, g, h, i)

	/*
	 * 定数
	 */
	const value float64 = 0.1

	// 列挙
	const (
		zero = iota
		one
		two
		three
	)
	fmt.Println(zero, one, two, three)

	// 型の検査
	fmt.Println("型の検査：", reflect.TypeOf(value))

	/*
		！！！変数・定数・関数定義の１文字目に注意！！！
		[スコープ]
		- 小文字から始まる場合、パッケージ内でのみ参照可能
		- 大文字から始まる場合、パッケージ外からも参照可能
			- `fmt.Println`
		[整数リテラル]
		- 8進数：01
		- 16進数：0x1
	*/
}
