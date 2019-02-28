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
	a, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	b, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	c, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	x, _ := strconv.Atoi(sc.Text())
	count := 0
	for i := 0; i <= a; i++ {
		for j := 0; j <= b; j++ {
			for k := 0; k <= c; k++ {
				if 500*i+100*j+50*k == x {
					count++
				}
			}
		}
	}
	return fmt.Sprintf("%d", count)
}

func main() {
	fmt.Println(exec(os.Stdin))
}
