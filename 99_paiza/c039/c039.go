package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func c039(r io.Reader) string {
	sc := bufio.NewScanner(r)
	sc.Scan()
	in := sc.Text()
	strs := strings.Split(in, "+")
	nums := make([]int, len(strs))
	for i, str := range strs {
		one := 0
		ten := 0
		for _, word := range str {
			if word == '<' {
				ten++
			} else if word == '/' {
				one++
			}
		}
		nums[i] = 10*ten + one
	}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return fmt.Sprintf("%d", sum)
}

func main() {
	fmt.Println(c039(os.Stdin))
}
