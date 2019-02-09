package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	/*
		sc.Scan()
		in, _ := strconv.Atoi(sc.Text())
		m := float64(in)
		sc.Scan()
		in, _ = strconv.Atoi(sc.Text())
		p := float64(in)
		sc.Scan()
		in, _ = strconv.Atoi(sc.Text())
		q := float64(in)
		fmt.Println(float64(m * (100 - p) / 100 * (100 - q) / 100))
	*/
	var in []float64
	for sc.Scan() {
		e, _ := strconv.Atoi(sc.Text())
		in = append(in, float64(e))
	}
	m, p, q := in[0], in[1], in[2]
	fmt.Println(float64(m * (100 - p) / 100 * (100 - q) / 100))
}
