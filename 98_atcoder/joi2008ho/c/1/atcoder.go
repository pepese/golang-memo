package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func exec(r io.Reader) string {
	sc := bufio.NewScanner(r)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	list := make([]string, n)
	col, num, white := "", 0, 0
	for i := 0; i < n; i++ {
		sc.Scan()
		in := sc.Text()
		if in == "0" {
			white++
		}
		if i == 0 {
			col, num = in, 1
		} else if col == in {
			num++
		}
		list[i] = in
		if (i+1)%2 == 0 && list[i-1] != in && in == "0" {
			white += num
			col = "0"
		} else if (i+1)%2 == 0 && list[i-1] != in && in == "1" {
			white -= num
			col = "1"
		} else if (i+1)%2 == 1 && col != in {
			col, num = in, 1
		}
		/*
			if (i+1)%2 == 0 && list[i-1] != in {
				tmp := list[i-1]
				c := 1
				for i-c >= 0 {
					if list[i-c] == tmp {
						list[i-c] = in
						c++
					} else {
						break
					}
				}
			}*/
	}
	/*
		c := 0
		for _, m := range list {
			if m == "0" {
				c++
			}
		}
		return fmt.Sprintf("%d", c)*/
	return fmt.Sprintf("%d", white)
}

func main() {
	fmt.Println(exec(os.Stdin))
}
