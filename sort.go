package main

import (
	"fmt"
)

func sort(list []int) {
	eNum := len(list)
	for i := eNum; i > 0; i-- {
		for j := 0; j < i-1; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
}

func main() {
	list := []int{3, 4, 1, 2, 8, 5}
	sort(list)
	fmt.Println(list)
}
