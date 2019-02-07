package main

// io Fscanf/Printf の型とかファイル入出力整理
// http://golang.jp/pkg/fmt
// https://text.baldanders.info/golang/string-and-rune/

import (
	"fmt"
	"os"
)

func main() {
	var a, b, c rune
	fmt.Fscanf(os.Stdin, "%d %d %d", &a, &b, &c)
	fmt.Printf("%d/%d/%d", a, b, c)
	//fmt.Println(strings.Replace(in, " ", "/", 2))
	/*
		sc := bufio.NewScanner(os.Stdin)
		sc.Scan()
		a := sc.Text()
		fmt.Println(a)
	*/
}
