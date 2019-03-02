package main

//https://qiita.com/drken/items/fd4e5e3630d0f5859067

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func search(s string, words []string, index int, result *string) {
	for _, w := range words {
		if index < len(s) && strings.Index(s[index:], w) == 0 {
			fmt.Printf("s:%s, index:%d, w:%s, now:%s\n", s, index, w, s[index:])
			if index+len(w) == len(s) {
				*result = "YES"
			} else {
				search(s, words, index+len(w), result)
			}
		}
	}
}

func search2(s string, words []string, result *string) {
	for _, w := range words {
		if strings.LastIndex(s, w) == len(s)-len(w) {
			fmt.Printf("s:%s, w:%s\n", s, w)
			if len(w) == len(s) {
				*result = "YES"
			} else if len(s) > len(w) {
				search2(s[0:len(s)-len(w)], words, result)
			}
		}
	}
}

func search3(s string, words []string) bool {
	flag := true
	for i := 0; i < len(s); {
		flag2 := false
		for _, w := range words {
			if strings.Index(s[i:], w) == 0 {
				flag2 = true
				i += len(w)
			}
		}
		if !flag2 {
			flag = false
			break
		}
	}
	return flag
}

func exec(r io.Reader) string {
	//words := []string{"dreamer", "dream", "eraser", "erase"}
	words := []string{"remaerd", "maerd", "resare", "esare"}
	sc := bufio.NewScanner(r)
	sc.Scan()
	s := sc.Text()
	/*
			index, result := 0, "NO"
			search(s, words, index, &result)
		return result
	*/
	/*
			result := "NO"
			search2(s, words, &result)
		return result
	*/
	s = func(s string) string {
		rs := []rune(s)
		for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
			rs[i], rs[j] = rs[j], rs[i]
		}
		return string(rs)
	}(s)
	if search3(s, words) {
		return "YES"
	} else {
		return "NO"
	}
}

func main() {
	fmt.Println(exec(os.Stdin))
}
