package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
type Comps [][]int

func (c Comps) Len() int {
	return len(c)
}

func (c Comps) Less(i, j int) bool {
	return c[i][3] < c[j][3]
}

func (c Comps) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}*/

func exec(r io.Reader) string {
	sc := bufio.NewScanner(r)
	sc.Scan()
	in := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(in[0])
	m, _ := strconv.Atoi(in[1])
	comps := make([][][]int, 101)
	for i := 0; i < n; i++ {
		sc.Scan()
		in := strings.Split(sc.Text(), " ")
		a, _ := strconv.Atoi(in[0])
		b, _ := strconv.Atoi(in[1])
		c, _ := strconv.Atoi(in[2])
		w, _ := strconv.Atoi(in[3])
		if comps[a] == nil {
			comps[a] = make([][]int, 101)
		}
		if comps[a][b] == nil {
			comps[a][b] = make([]int, 101)
		}
		if comps[a][b][c] < w {
			comps[a][b][c] = w
		}
	}
	for i := 0; i < 101; i++ {
		if comps[i] == nil {
			comps[i] = make([][]int, 101)
		}
		for j := 0; j < 101; j++ {
			if comps[i][j] == nil {
				comps[i][j] = make([]int, 101)
			}
			for k := 0; k < 101; k++ {
				if i > 1 && comps[i][j][k] < comps[i-1][j][k] {
					comps[i][j][k] = comps[i-1][j][k]
				}
				if j > 1 && comps[i][j][k] < comps[i][j-1][k] {
					comps[i][j][k] = comps[i][j-1][k]
				}
				if k > 1 && comps[i][j][k] < comps[i][j][k-1] {
					comps[i][j][k] = comps[i][j][k-1]
				}
			}
		}
	}
	result := ""
	for i := 0; i < m; i++ {
		sc.Scan()
		in := strings.Split(sc.Text(), " ")
		x, _ := strconv.Atoi(in[0])
		y, _ := strconv.Atoi(in[1])
		z, _ := strconv.Atoi(in[2])
		result += fmt.Sprintf("%d\n", comps[x][y][z])
	}
	/*
		comps := make([][]int, n)
		usrs := make([][]int, m)
		result := ""
		for i := 0; i < n; i++ {
			sc.Scan()
			in := strings.Split(sc.Text(), " ")
			a, _ := strconv.Atoi(in[0])
			b, _ := strconv.Atoi(in[1])
			c, _ := strconv.Atoi(in[2])
			w, _ := strconv.Atoi(in[3])
			comps[i] = []int{a, b, c, w}
		}
		sort.Sort(sort.Reverse(Comps(comps)))
		for i := 0; i < m; i++ {
			sc.Scan()
			in := strings.Split(sc.Text(), " ")
			x, _ := strconv.Atoi(in[0])
			y, _ := strconv.Atoi(in[1])
			z, _ := strconv.Atoi(in[2])
			usrs[i] = []int{x, y, z}
			for j := 0; j < n; j++ {
				if usrs[i][0] >= comps[j][0] && usrs[i][1] >= comps[j][1] && usrs[i][2] >= comps[j][2] {
					result += fmt.Sprintf("%d\n", comps[j][3])
					break
				} else if j == n-1 {
					result += "0\n"
				}
			}
		}
	*/
	return result[0 : len(result)-1]
}

func main() {
	fmt.Println(exec(os.Stdin))
}
