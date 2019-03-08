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
	m, _ := strconv.Atoi(in[1])
	N := make([]int, n)
	for i := 0; i < m; i++ {
		sc.Scan()
		in = strings.Split(sc.Text(), " ")
		a, _ := strconv.Atoi(in[0])
		b, _ := strconv.Atoi(in[1])
		N[b]++
	}
	return ""
}

func main() {
	fmt.Println(exec(os.Stdin))
}
