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
	//n, _ := strconv.Atoi(sc.Text())
	var results []string
	for sc.Scan() {
		var even, odd int
		str := sc.Text()
		for i, v := range str {
			n := string(v)
			if n == "X" { // ムシ
			} else if i%2 == 0 { // 偶数桁
				num, _ := strconv.Atoi(n)
				tmp := 2 * num
				if tmp >= 10 {
					tmp = tmp/10 + tmp%10
				}
				even += tmp
			} else { // 奇数桁
				num, _ := strconv.Atoi(n)
				odd += num
			}
		}
		X := 10 - (even+odd)%10
		if X == 10 {
			X = 0
		}
		results = append(results, fmt.Sprintf("%d", X))
	}
	return strings.Join(results, "\n")
}

func main() {
	fmt.Println(exec(os.Stdin))
}
