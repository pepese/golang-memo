package main

import "fmt"

func delete(list []int, i int) []int {
	//fmt.Println(list, i)
	result := list[:i]
	result = append(result, list[i+1:]...)
	return result
}

func permutation(list []int) [][]int {
	var result [][]int
	var inner func(in, out []int)
	inner = func(in, out []int) {
		if len(in) == 1 {
			//out = append(out, in[0])
			result = append(result, append(out, in[0]))
		} else {
			for i, n := range in {
				//out = append(out, n)
				//in = delete(in, i)
				inner(delete(in, i), append(out, n))
			}
		}
	}
	inner(list, []int{})
	return result
}

func main() {
	fmt.Println(permutation([]int{1, 2, 3}))
	//fmt.Println(delete([]int{1, 2, 3}, 0))
}
