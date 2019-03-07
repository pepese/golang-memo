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
	in := strings.Split(sc.Text(), " ")
	R, _ := strconv.Atoi(in[0])
	C, _ := strconv.Atoi(in[1])
	arr := make([][]int, R)
	var sx, sy, tx, ty int
	for i := 0; i < R; i++ {
		sc.Scan()
		in := strings.Split(sc.Text(), "")
		arr[i] = make([]int, C)
		for j, arj := range in {
			if arj == "s" {
				sx, sy = j, i
			} else if arj == "t" {
				tx, ty = j, i
			} else {
				arr[i][j], _ = strconv.Atoi(arj)
			}
		}
	}
	result := math.MaxInt64
	var search func(nowx, nowy, called int)
	search = func(nowx, nowy, called int) {
		if nowx == tx && nowy == ty {
			if result > called {
				result = called
			}
			return
		} else if result <= called {
			return
		} else {
			if nowx < tx {
				search(nowx+1, nowy, called+arr[nowy][nowx+1])
			}
			if nowy < ty {
				if nowy%2 == 0 {
					search(nowx, nowy+1, called+arr[nowy+1][nowx])
				} else if nowy%2 == 1 && nowx < tx {
					search(nowx+1, nowy+1, called+arr[nowy+1][nowx+1])
				} else {
					return
				}
			}
		}
	}
	search(sx, sy, 0)
	return fmt.Sprintf("%d", result)
}

func main() {
	fmt.Println(exec(os.Stdin))
}
