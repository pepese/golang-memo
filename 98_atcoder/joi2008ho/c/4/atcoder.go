package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func exec(r io.Reader) string {
	sc := bufio.NewScanner(r)
	sc.Scan()
	in := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(in[0])
	m, _ := strconv.Atoi(in[1])
	list := make([]map[int]int, n)
	maxs := make()
	for i := 0; i < n; i++ {
		sc.Scan()
		in := strings.Split(sc.Text(), " ")
		num, _ := strconv.Atoi(in[0])
		column := make(map[int]int, num)
		sube := 1000 * 2 * 10 * 2
		for j := 1; j <= num; j++ {
			tmp1, _ := strconv.Atoi(in[2*j-1])
			tmp2, _ := strconv.Atoi(in[2*j])
			column[tmp1] = tmp2
		}
		list[i] = column
	}
	sort.Sort(sort.IntSlice(list))
	list = list[:n-m]
	result := 0
	for _, mm := range list {
		result += mm
	}
	return fmt.Sprintf("%d", result)
}

func main() {
	fmt.Println(exec(os.Stdin))
}
