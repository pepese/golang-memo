package main

import "fmt"

func binarySearch(list []int, elem int) int {
	start, end, index := 0, len(list)-1, 0
	for true {
		index = (start + end) / 2
		if end < start {
			return -1
		} else if list[index] == elem {
			return index
		} else if list[index] < elem {
			start = index + 1
		} else if list[index] > elem {
			end = index - 1
		}
	}
	return -1
}

func main() {
	list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(binarySearch(list, 6))
}
