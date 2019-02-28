package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
)

func exec(r io.Reader) string {
	sc := bufio.NewScanner(r)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	list := make([]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		list[i], _ = strconv.Atoi(sc.Text())
	}
	sort.Sort(sort.IntSlice(list))
	c, tmp := 0, -1
	for _, m := range list {
		if tmp != m {
			tmp = m
			c++
		}
	}
	return fmt.Sprintf("%d", c)
}

func main() {
	fmt.Println(exec(os.Stdin))
}
