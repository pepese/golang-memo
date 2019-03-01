package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var (
	yes = "YES"
	no  = "NO"
)

func search(s string, words []string, index int, result *string) {
	for _, w := range words {
		if index < len(s) && strings.Index(s[index:], w) == 0 {
			if index+len(w) == len(s) {
				result = &yes
			} else {
				search(s, words, index+len(w), result)
			}
		}
		fmt.Printf("%s, w: %s, index: %d, result: %s\n", s[index:], w, strings.Index(s[index:], w), *result)
	}
}

func exec(r io.Reader) string {
	words := []string{"dreamer", "dream", "eraser", "erase"}
	sc := bufio.NewScanner(r)
	sc.Scan()
	s := sc.Text()
	/*
		index, result := 0, "NO"
		for true {
			flag := false
			for i, w := range words {
				if index < len(s) && strings.Index(s[index:], w) == 0 {
					if index+len(w) == len(s) {
						result = "YES"
					}
					if i == 0 && (strings.Index(s[index+5:], words[2]) == 0 || strings.Index(s[index+5:], words[3]) == 0) {
						continue
					} else {
						index += len(w)
						break
					}
				} else if i == 3 && !flag {
					flag = true
				}
			}
			if result == "YES" || flag {
				break
			}
		}
		return result
	*/
	index := 0
	var result *string
	result = &no
	search(s, words, index, result)
	return *result
}

func main() {
	fmt.Println(exec(os.Stdin))
}
