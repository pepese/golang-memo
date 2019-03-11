package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func exec(r io.Reader) string {
	sc := bufio.NewScanner(r)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	d := make([][]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		in := strings.Split(sc.Text(), " ")
		d[i] = make([]int, n)
		for j, dij := range in {
			d[i][j], _ = strconv.Atoi(dij)
		}
	}
	sc.Scan()
	m, _ := strconv.Atoi(sc.Text())
	st := make([][]int, m)
	for i := 0; i < m; i++ {
		sc.Scan()
		in := strings.Split(sc.Text(), " ")
		si, _ := strconv.Atoi(in[0])
		ti, _ := strconv.Atoi(in[1])
		st[i] = []int{si, ti}
	}
	stCheck := func(si, score int) bool {
		for i := 0; i < len(st); i++ {
			if si == st[i][0] {
				if st[i][1] >= score {
					return true
				}
			}
		}
		return false
	}
	minTime := math.MaxInt64
	var search func(now, score int, memo []int)
	search = func(now, score int, memo []int) {
		for i := 0; i < len(memo); i++ {
			if memo[i] == 0 {
				memo[i] = 1
				nextScore := score + d[now][i]
				if stCheck(i, nextScore) && minTime > nextScore {
					fmt.Printf("%d to %d\n", now, i)
					minTime = nextScore
					search(i, nextScore, memo)
				}
			}
		}
	}
	search(0, 0, make([]int, n))
	return fmt.Sprintf("%d", minTime)
}

func main() {
	fmt.Println(exec(os.Stdin))
}
