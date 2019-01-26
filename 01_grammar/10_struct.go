package main

import (
	"fmt"
)

type user struct {
	name  string // var が不要
	score int
}

// レシーバで構造体とメソッド(関数)を紐付け。（厳密には「値レシーバ」）
// ちなみになぜ「レシーバ」と呼ぶのかというと、
// 昔の OOP 言語の文脈ではメソッド呼び出しのことを「メッセージの送信」と言い、
// メソッドを呼び出される側は「メッセージの受信側」だから。
func (u user) show() {
	fmt.Printf("name: %s, socre: %d\n", u.name, u.score)
}

// 上記のメソッド定義ではuserは値渡し(コピー)でメソッドに渡される。
// *をつける事で参照渡しになり、「ポインタレシーバ」という。
// 値レシーバではメモリの無駄なので、基本的にはポインタレシーバでメソッド定義する。
func (u *user) hit() {
	u.score++
}

// 参照渡しでないと
func (u user) hit2() {
	u.score++
}

func main() {
	// ポインタを返す
	u1 := new(user) // ちなみに「 u1 := &user{name: "pepese", score: 100"}」でも同じ意味。
	u1.name = "pepese"
	u1.score = 100
	fmt.Println(u1) // &{pepese 100}
	(*u1).name = "hyperse"
	(*u1).score = 1000
	// ポインタの場合は、 (*u1) と書くべきだが、 u1 でもコンパイラが解釈してくれる。
	fmt.Println(u1) // &{hyperse 1000}
	// メソッドの利用
	u1.hit2()
	u1.show() // name: hyperse, socre: 1000
	u1.hit()
	u1.show() // name: hyperse, socre: 1001

	// 構造体そのもの（値）を返す、ポインタでないことに注意
	u2 := user{name: "pepese2", score: 200} // 「&」がついてないことに注意！
	fmt.Println(u2)                         // {pepese2 200}
	u2.name = "hyperse"                     // {hyperse 200}
	// 構造体を返却しているので、 (*u2).name とするとエラーになる。
	fmt.Println(u2)
	// メソッドの利用
	u2.hit2()
	u2.show() // name: hyperse, socre: 200
	u2.hit()
	u2.show() // name: hyperse, socre: 201
}

// レシーバの神解説
// https://skatsuta.github.io/2015/12/29/value-receiver-pointer-receiver/

// tag
// https://qiita.com/itkr/items/9b4e8d8c6d574137443c

// enum
// https://qiita.com/awakia/items/c81c7816b9aea5422250
