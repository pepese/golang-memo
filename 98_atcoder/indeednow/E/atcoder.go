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
	n, _ := strconv.Atoi(in[0])
	m, _ := strconv.Atoi(in[1])
	N := make([]int, n)
	mp := make([][]int, n)
	for i := 0; i < m; i++ {
		sc.Scan()
		in = strings.Split(sc.Text(), " ")
		a, _ := strconv.Atoi(in[0])
		b, _ := strconv.Atoi(in[1])
		if mp[b-1] == nil {
			mp[b-1] = []int{a - 1}
		} else {
			mp[b-1] = append(mp[b-1], a-1)
		}
		N[b-1]++
	}
	X := make([][]float64, n)
	for i := 0; i < n; i++ {
		X[i] = make([]float64, n+1)
		X[i][i], X[i][n] = 1.0, 0.1
		for _, j := range mp[i] {
			if N[i] != 0 {
				X[i][j] = -0.9 / float64(N[j])
			}
		}
	}
	// ガウスの消去法（掃出し法）: X[n][n+1]float64
	for i := 0; i < n; i++ {
		for j := 0; j < n+1; j++ { // i 行目を [i][i] で割って [i][i] を 1 にする
			if i != j {
				X[i][j] /= X[i][i]
			}
		}
		X[i][i] = 1
		for j := 0; j < n; j++ { // i 行目以外の j 行目 の i 列要素が 0 になるように引く
			if i != j {
				tmp := X[j][i]
				for k := 0; k < n+1; k++ {
					X[j][k] -= tmp * X[i][k]
				}
			}
		}
	}
	result := ""
	for i := 0; i < n; i++ {
		result += fmt.Sprintf("%0.16f\n", X[i][n])
	}
	return result
}

func main() {
	fmt.Println(exec(os.Stdin))

}
