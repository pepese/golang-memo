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
	N, _ := strconv.Atoi(in[0])
	Y, _ := strconv.Atoi(in[1])
	result := ""
	sx := Y / 10000
	sy := (Y - 10000*sx) / 5000
	for x := sx; x >= 0; x-- { // 10000
		for y := sy + 2*x; y >= 0; y-- { // 5000
			z := N - x - y // 1000
			if z >= 0 && 10000*x+5000*y+1000*z == Y {
				result = fmt.Sprintf("%d %d %d", x, y, z)
				break
			}
		}
		if result != "" {
			break
		}
	}
	if result == "" {
		result = "-1 -1 -1"
	}
	return result
}

func main() {
	fmt.Println(exec(os.Stdin))
}
