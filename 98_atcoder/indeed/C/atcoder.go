package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func exec(r io.Reader) string {
	sc := bufio.NewScanner(r)
	sc.Scan()
	in := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(in[0])
	l, _ := strconv.Atoi(in[1])
	arr := make([]int, n)
	sc.Scan()
	in = strings.Split(sc.Text(), " ")
	for i := 0; i < n; i++ {
		ai, _ := strconv.Atoi(in[i])
		arr[i] = ai
	}
	maxi := -1
	result := ""
	for s := 0; s < n-l+1; s++ {
		if l == 1 {
			result += fmt.Sprintf("%d\n", arr[s])
		} else {
			if maxi < s {
				max := -1000000000
				for i := s; i < s+l; i++ {
					if max < arr[i] {
						max = arr[i]
						maxi = i
					}
				}
			} else if arr[s+l-1] > arr[maxi] {
				maxi = s + l - 1
			}
			result += fmt.Sprintf("%d\n", arr[maxi])
		}
	}
	return result[:len(result)-1]
}

func main() {
	fmt.Println(exec(os.Stdin))
}
