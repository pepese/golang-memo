package main

import (
	"fmt"
)

func sample(pa *int) {
	fmt.Println("---func sample start")
	fmt.Println(pa)
	fmt.Println(*pa)
	fmt.Println(&pa) // ポインタを格納しているアドレスは main と変わっている
	fmt.Println("---change value")
	*pa = 3
	fmt.Println(pa)
	fmt.Println(*pa)
	fmt.Println(&pa)
	fmt.Println("---func sample end")
}

func main() {
	a := 5

	var pa *int
	pa = &a

	fmt.Println(pa)  // ポインタが参照しているアドレス：0xc000016078
	fmt.Println(*pa) // アドレスに格納している値：5
	fmt.Println(&a)  // a のアドレス（paと同じ）：0xc000016078
	fmt.Println(&pa) // ポインタを格納しているアドレス：0xc00000c028

	fmt.Println("---func sample call")
	sample(pa)

	fmt.Println(pa)
	fmt.Println(*pa)
	fmt.Println(&pa)
}
