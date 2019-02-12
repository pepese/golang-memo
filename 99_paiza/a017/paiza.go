package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// 作成ロジック用関数
func exec(r io.Reader) string {
	sc := bufio.NewScanner(r)
	sc.Scan()
	ins := strings.Split(sc.Text(), " ")
	h, _ := strconv.Atoi(ins[0])
	w, _ := strconv.Atoi(ins[1])
	//n, _ := strconv.Atoi(ins[2])
	results := make([][]string, h)
	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			results[i] = append(results[i], ".")
		}
	}
	for sc.Scan() {
		ins = strings.Split(sc.Text(), " ")
		hi, _ := strconv.Atoi(ins[0])
		wi, _ := strconv.Atoi(ins[1])
		xi, _ := strconv.Atoi(ins[2])
		flag := true
		y := 0
		for i := 0; i < h && flag; i++ { // 底にあたる高さを調べる
			for j := xi; j < xi+wi && flag; j++ {
				if results[i][j] == "#" {
					flag = false
				}
			}
			if flag {
				y = i
			}
		}
		for i := y - hi + 1; i <= y; i++ {
			for j := xi; j < xi+wi; j++ {
				results[i][j] = "#"
			}
		}
	}

	var tmp []string
	for i := 0; i < h; i++ {
		tmp = append(tmp, strings.Join(results[i], ""))
	}
	return strings.Join(tmp, "\n") // 結果の文字列を返却
}

// 標準入力は `ctr+D` で EOF になる
func main() {
	fmt.Println(exec(os.Stdin))
}
