package main

import (
	"fmt"
)

type user struct {
	name  string // var が不要
	score int
}

// レシーバで構造体とメソッド(関数)を紐付け
func (u user) show() {
	fmt.Printf("name: %s, socre: %d\n", u.name, u.score)
}

// デフォルトではuserには値渡し(コピー)でメソッドに渡される。
// *をつける事で参照渡しになる
func (u *user) hit() {
	u.score++
}

// 参照渡しでないと
func (u user) hit2() {
	u.score++
}

func main() {
	// ポインタを返す
	u1 := new(user)
	u1.name = "pepese"
	u1.score = 100
	fmt.Println(u1)
	(*u1).name = "hyperse"
	(*u1).score = 1000
	fmt.Println(u1)

	// 構造体を返す
	u2 := user{name: "pepese2", score: 200}
	fmt.Println(u2)
	u2.name = "hyperse"
	fmt.Println(u2)

	// メソッドの利用
	u1.hit2()
	u1.show()
	u1.hit()
	u1.show()
}

// tag
// https://qiita.com/itkr/items/9b4e8d8c6d574137443c

// enum
// https://qiita.com/awakia/items/c81c7816b9aea5422250
