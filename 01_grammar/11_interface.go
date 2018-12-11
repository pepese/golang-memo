package main

import (
	"fmt"
)

type greeter interface {
	greet() // func が不要
}

type japanese struct {
	end string
}
type english struct {
	end string
}

func (ja japanese) greet() {
	fmt.Println("こんにちは" + ja.end)
}

func (en english) greet() {
	fmt.Println("Hello " + en.end)
}

// 空のインタフェース型で引数を受け取る
func show(t interface{}) {
	// 型アサーション
	// 2つの値が返る
	_, ok := t.(japanese)
	// okを使って条件分岐
	if ok {
		fmt.Println("i am japanese")
	} else {
		fmt.Println("i am not japan")
	}

	// 型Switch
	switch t.(type) {
	case japanese:
		fmt.Println("僕は日本人だよ")
	default:
		fmt.Println("僕は日本人じゃないよ")
	}
}

func main() {
	ja := japanese{end: "^^"}
	en := english{end: ":)"}
	var greeter1 greeter
	var greeter2 greeter
	greeter1 = ja
	greeter2 = en
	greeter1.greet()
	greeter2.greet()

	// 空のインターフェースの確認
	greeters := []greeter{japanese{}, english{}}
	for _, greeter := range greeters {
		greeter.greet()
		show(greeter)
	}
}
