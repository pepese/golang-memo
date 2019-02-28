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
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	in := strings.Split(sc.Text(), " ")
	a := make([]int, n)
	for i, m := range in {
		a[i], _ = strconv.Atoi(m)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	var alice, bob int
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			alice += a[i]
		} else {
			bob += a[i]
		}
	}
	return fmt.Sprintf("%d", alice-bob)
}

func main() {
	fmt.Println(exec(os.Stdin))
}
