package main

// https://qiita.com/drken/items/7c6ff2aa4d8fce1c9361#bit-%E5%85%A8%E6%8E%A2%E7%B4%A2f

import "fmt"

func main() {
	var n uint = 2                        // bit の桁数
	for bit := 0; bit < (1 << n); bit++ { // n ビットまでループする
		fmt.Printf("%b\n", bit)
	}
}
