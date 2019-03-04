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
	sc.Scan()
	in := strings.Split(sc.Text(), " ")
	list := make([]int, len(in))
	for i, n := range in {
		list[i], _ = strconv.Atoi(n)
	}
	sort.Sort(sort.IntSlice(list))
	max := list[len(list)-1]
	c := 0
	for _, n := range list[:len(list)-1] {
		c += n
	}
	if max < c {
		return "Yes"
	} else {
		return "No"
	}
}

func main() {
	fmt.Println(exec(os.Stdin))
}
