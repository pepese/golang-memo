package main

import "fmt"

func main() {
	var n uint = 2                        // ビットの桁数
	for bit := 0; bit < (1 << n); bit++ { // 1 から n ビットまでの組合せ
		fmt.Printf("%b\n", bit)
	}
	fmt.Println("---")
	for bit := 1 << (n - 1); bit < (1 << n); bit++ { // n ビットの組合せ
		fmt.Printf("%b", bit)
		if (bit & (1 << (2 - 1))) > 0 { // 2 ビット目が立っているか確認
			fmt.Println(", 2nd bit flaged.")
		} else {
			fmt.Println()
		}
	}
}
