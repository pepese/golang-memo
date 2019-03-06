package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func delEle(list [][]int, i int) [][]int {
	result := [][]int{}
	result = append(result, list[:i]...)
	result = append(result, list[i+1:]...)
	return result
}

func permutation(list [][]int) [][][]int {
	var result [][][]int
	var inner func(in, out [][]int)
	inner = func(in, out [][]int) {
		if len(in) == 1 {
			result = append(result, append(out, in[0]))
		} else {
			for i, n := range in {
				inner(delEle(in, i), append(out, n))
			}
		}
	}
	inner(list, [][]int{})
	return result
}

func exec(r io.Reader) string {
	chikan := func(s, p []int, nn int) []int {
		tmp := make([]int, nn)
		for i := 0; i < nn; i++ {
			si := s[i]
			pi := p[i] - 1
			tmp[pi] = si
		}
		return tmp
	}
	hikaku := func(r, t []int) bool { // r > t なら true
		for i := 0; i < len(r); i++ {
			if r[i] < t[i] {
				return false
			} else if r[i] > t[i] {
				return true
			}
		}
		return true
	}
	sc := bufio.NewScanner(r)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	sStr := strings.Split(sc.Text(), " ")
	var S []int
	for i := range sStr {
		num, _ := strconv.Atoi(sStr[i])
		S = append(S, num)
	}
	sc.Scan()
	m, _ := strconv.Atoi(sc.Text())
	p := make([][]int, m)
	for i := 0; sc.Scan(); i++ {
		pStr := strings.Split(sc.Text(), " ")
		p[i] = make([]int, n)
		for j := range pStr {
			num, _ := strconv.Atoi(pStr[j])
			p[i][j] = num
		}
	}
	//fmt.Println(permutation(p))
	result := S
	for _, ii := range permutation(p) {
		var bit uint
		for bit = 1 << uint(len(ii)-1); bit < (1 << uint(len(ii))); bit++ {
			T := S
			for j := 0; j < len(ii); j++ {
				if (bit & (1 << uint(j))) > 0 {
					T = chikan(T, ii[j], n)
				}
			}
			flag := hikaku(result, T)
			if flag {
				result = T
			}
		}
	}
	var sr []string
	for _, num := range result {
		str := fmt.Sprintf("%d", num)
		sr = append(sr, str)
	}
	return strings.Join(sr, " ")

	/*
		for i := 0; sc.Scan(); i++ {
			pStr := strings.Split(sc.Text(), " ")
			p[i] = make([]int, n)
			for j := range pStr {
				num, _ := strconv.Atoi(pStr[j])
				p[i][j] = num
			}
		}
		pp := func(xs []int) {
			T := S
			for j := 0; j < len(xs); j++ {
				for i := 0; i <= j; i++ {
					T = chikan(T, p[xs[i]-1], n)
				}
				flag := hikaku(result, T)
				if flag {
					result = T
				}
			}
		}
		permutation(pp, m, m)
		var sr []string
		for _, num := range result {
			str := fmt.Sprintf("%d", num)
			sr = append(sr, str)
		}
		return strings.Join(sr, " ")
	*/
}

func main() {
	fmt.Println(exec(os.Stdin))
}
