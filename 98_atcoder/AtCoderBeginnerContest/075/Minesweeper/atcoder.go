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
	h, _ := strconv.Atoi(in[0])
	w, _ := strconv.Atoi(in[1])
	hw := make([][]string, h)
	for i := 0; i < h; i++ {
		sc.Scan()
		in = strings.Split(sc.Text(), "")
		hw[i] = in
	}
	for i := range hw {
		for j := range hw[i] {
			if hw[i][j] == "." {
				c := 0
				for y := -1; y < 2; y++ {
					for x := -1; x < 2; x++ {
						if x == 0 && y == 0 {
							continue
						}
						yy, xx := i+y, j+x
						if xx >= 0 && yy >= 0 && xx < w && yy < h {
							if hw[yy][xx] == "#" {
								c++
							}
						}
					}
				}
				hw[i][j] = fmt.Sprintf("%d", c)
			}
		}
	}
	result := func(hw [][]string) string {
		str := ""
		for _, line := range hw {
			str += fmt.Sprintf("%s\n", strings.Join(line, ""))
		}
		return str[0 : len(str)-1]
	}(hw)
	return result
}

func main() {
	fmt.Println(exec(os.Stdin))
}
