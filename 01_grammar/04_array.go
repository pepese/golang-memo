package main

import (
	"fmt"
)

func main() {
	var a [3]int
	a[0] = 0
	a[1] = 1
	a[2] = 2
	fmt.Println(a) // [0 1 2]

	b := [3]int{3, 4, 5}
	fmt.Println(b) // [3 4 5]

	c := [...]int{6, 7, 8}
	fmt.Println(c) // [6 7 8]
}
