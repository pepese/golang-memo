package main

import (
	"fmt"
)

func main() {
	var a [3]int
	a[0] = 0
	a[1] = 1
	a[2] = 2
	fmt.Println(a)

	b := [3]int{3, 4, 5}
	fmt.Println(b)

	c := [...]int{6, 7, 8}
	fmt.Println(c)
}
