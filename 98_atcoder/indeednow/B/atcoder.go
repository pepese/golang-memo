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
	//start := time.Now()
	sc := bufio.NewScanner(r)
	sc.Scan()
	in := strings.Split(sc.Text(), " ")
	R, _ := strconv.Atoi(in[0])
	C, _ := strconv.Atoi(in[1])
	arr := make([][]int, R)
	memos := make([][]int, R)
	var sx, sy, tx, ty int
	for i := 0; i < R; i++ {
		sc.Scan()
		in := strings.Split(sc.Text(), "")
		arr[i] = make([]int, C)
		memos[i] = make([]int, C)
		for j, arj := range in {
			if arj == "s" {
				sx, sy = j, i
			} else if arj == "t" {
				tx, ty = j, i
			} else {
				arr[i][j], _ = strconv.Atoi(arj)
			}
			memos[i][j] = 99 * 99 * 9
		}
	}
	check := func(nowx, nowy, tox, toy, score int) (int, bool) {
		if tox < 0 || tox >= C || toy < 0 || toy >= R || (tox == sx && toy == sy) {
			return score, false
		}
		if nowx == sx && nowy == sy {
			memos[toy][tox] = arr[toy][tox]
			return memos[toy][tox], true
		} else if memos[toy][tox] > arr[toy][tox]+score {
			memos[toy][tox] = arr[toy][tox] + score
			return memos[toy][tox], true
		}
		return score, false
	}
	var search func(nowx, nowy, score int)
	search = func(nowx, nowy, score int) {
		//now := time.Now()
		if memos[ty][tx] <= score /*|| now.Add(-1900*time.Millisecond).After(start) */ {
			return
		}
		if !(nowx == tx && nowy == ty) {
			if num, flag := check(nowx, nowy, nowx-1, nowy, score); flag { // 左
				search(nowx-1, nowy, num)
			}
			if num, flag := check(nowx, nowy, nowx+1, nowy, score); flag { // 右
				search(nowx+1, nowy, num)
			}
			if nowy%2 == 0 {
				if num, flag := check(nowx, nowy, nowx-1, nowy-1, score); flag { // 左上
					search(nowx-1, nowy-1, num)
				}
				if num, flag := check(nowx, nowy, nowx, nowy-1, score); flag { // 右上
					search(nowx, nowy-1, num)
				}
				if num, flag := check(nowx, nowy, nowx-1, nowy+1, score); flag { // 左下
					search(nowx-1, nowy+1, num)
				}
				if num, flag := check(nowx, nowy, nowx, nowy+1, score); flag { // 右下
					search(nowx, nowy+1, num)
				}
			} else {
				if num, flag := check(nowx, nowy, nowx, nowy-1, score); flag { // 左上
					search(nowx, nowy-1, num)
				}
				if num, flag := check(nowx, nowy, nowx+1, nowy-1, score); flag { // 右上
					search(nowx+1, nowy-1, num)
				}
				if num, flag := check(nowx, nowy, nowx, nowy+1, score); flag { // 左下
					search(nowx, nowy+1, num)
				}
				if num, flag := check(nowx, nowy, nowx+1, nowy+1, score); flag { // 右下
					search(nowx+1, nowy+1, num)
				}
			}
		}
	}
	search(sx, sy, 0)
	return fmt.Sprintf("%d", memos[ty][tx])
}

func main() {
	fmt.Println(exec(os.Stdin))
}
