package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// 作成ロジック用関数
func exec(r io.Reader) string {
	sc := bufio.NewScanner(r)
	// ここからロジック（繋げるだけ）
	var str []string
	for sc.Scan() {
		str = append(str, fmt.Sprint(sc.Text()))
	}
	return strings.Join(str, "\n") // 結果の文字列を返却
}

// 標準入力は `ctr+D` で EOF になる
func main() {
	fmt.Println(exec(os.Stdin))
}
