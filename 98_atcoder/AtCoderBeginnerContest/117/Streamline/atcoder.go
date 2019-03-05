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

type interval struct {
	from, to, val int
}

type intervals []interval

func (p intervals) Len() int {
	return len(p)
}

func (p intervals) Less(i, j int) bool {
	return p[i].val < p[j].val
}

func (p intervals) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type intervals2 []interval

func (p intervals2) Len() int {
	return len(p)
}

func (p intervals2) Less(i, j int) bool {
	return p[i].from < p[j].from
}

func (p intervals2) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func exec(r io.Reader) string {
	sc := bufio.NewScanner(r)
	sc.Scan()
	in := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(in[0])
	m, _ := strconv.Atoi(in[1])
	if n >= m {
		return "0"
	}
	sc.Scan()
	in = strings.Split(sc.Text(), " ")
	list := make([]int, m)
	for i, x := range in {
		num, _ := strconv.Atoi(x)
		list[i] = num
	}
	var ivals intervals
	sort.Sort(sort.IntSlice(list))
	for i := 0; i < len(list)-1; i++ {
		ii := interval{from: i, to: i + 1, val: list[i+1] - list[i]}
		ivals = append(ivals, ii)
		sort.Sort(sort.Reverse(ivals))
		if len(ivals) > n-1 {
			ivals = ivals[:n-1]
		}
	}
	ivals2 := intervals2(ivals)
	sort.Sort(ivals2)
	start, count := ivals2[0].to, list[ivals2[0].from]-list[0]
	for i := 1; i < len(ivals2); i++ {
		diff := list[ivals2[i].from] - list[start]
		count += diff
		start = ivals2[i].to
	}
	count += list[len(list)-1] - list[ivals2[n-2].to]
	return fmt.Sprintf("%d", count)
}

func main() {
	fmt.Println(exec(os.Stdin))
}
