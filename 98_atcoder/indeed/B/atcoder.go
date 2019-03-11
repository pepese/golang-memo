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
	rooms := make([]int, n)
	roomCheck := func(c, d int) bool {
		for i := c - 1; i < d; i++ {
			if rooms[i] == 1 {
				return false
			}
		}
		return true
	}
	roomOccupied := func(c, d int) bool {
		for i := c - 1; i < d; i++ {
			rooms[i] = 1
		}
		return true
	}
	abi := make([][]int, m)
	for i := 0; i < m; i++ {
		sc.Scan()
		in = strings.Split(sc.Text(), " ")
		ai, _ := strconv.Atoi(in[0])
		bi, _ := strconv.Atoi(in[1])
		abi[i] = []int{ai, bi}
	}
	fmt.Println(abi)
	result := ""
	sc.Scan()
	q, _ := strconv.Atoi(sc.Text())
	for i := 0; i < q; i++ {
		sc.Scan()
		in = strings.Split(sc.Text(), " ")
		ci, _ := strconv.Atoi(in[0])
		di, _ := strconv.Atoi(in[1])
		if i >= m {
			if roomCheck(ci, di) {
				roomOccupied(ci, di)
				result += "OK\n"
			} else {
				result += "NG\n"
			}
		} else {
			if (abi[i][0] > di || abi[i][1] < ci) && roomCheck(ci, di) {
				result += "OK\n"
				roomOccupied(ci, di)
			} else if (abi[i][0] <= ci || ci <= abi[i][1]) || (abi[i][0] <= di || di <= abi[i][1]) {
				result += "NG\n"
			}
		}
	}
	return result[:len(result)-1]
}

func main() {
	fmt.Println(exec(os.Stdin))
}
