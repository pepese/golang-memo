package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	line := strings.Split(sc.Text(), " ")
	boxNum, _ := strconv.Atoi(line[0])
	r, _ := strconv.Atoi(line[1])
	rr := 2 * r
	var oks []int
	for i := 1; i <= boxNum; i++ {
		sc.Scan()
		line := strings.Split(sc.Text(), " ")
		h, _ := strconv.Atoi(line[0])
		w, _ := strconv.Atoi(line[1])
		d, _ := strconv.Atoi(line[2])
		if rr <= h && rr <= w && rr <= d {
			oks = append(oks, i)
		}
	}
	for _, ok := range oks {
		fmt.Println(ok)
	}
}
