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
	k, _ := strconv.Atoi(in[1])
	m, _ := strconv.Atoi(in[2])
	result := ""
	for i := 0; i < n; i++ {
		sc.Scan()
		page := i/k + 1
		if page == m {
			result += fmt.Sprintf("%s\n", sc.Text())
		} else if page > m {
			break
		}
	}
	return result[:len(result)-1]
}

func main() {
	fmt.Println(exec(os.Stdin))
}
