package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func exec(r io.Reader) string {
	sc := bufio.NewScanner(r)
	sc.Scan()
	//n, _ := strconv.Atoi(sc.Text())
	t, x, y := 0, 0, 0
	for sc.Scan() {
		in := strings.Split(sc.Text(), " ")
		ti, _ := strconv.Atoi(in[0])
		xi, _ := strconv.Atoi(in[1])
		yi, _ := strconv.Atoi(in[2])
		tt := int(math.Abs(float64(ti - t)))
		xx := int(math.Abs(float64(xi - x)))
		yy := int(math.Abs(float64(yi - y)))
		if (tt-xx-yy) >= 0 && (tt-xx-yy)%2 == 0 {
			t, x, y = ti, xi, yi
		} else {
			return "NO"
		}
	}
	return "YES"
}

func main() {
	fmt.Println(exec(os.Stdin))
}
