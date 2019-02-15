package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func member(n int, xs []int) bool {
	for _, x := range xs {
		if n == x {
			return true
		}
	}
	return false
}

func permSub(f func([]int), n, m int, xs []int) {
	if len(xs) == m {
		f(xs)
	} else {
		for i := 1; i <= n; i++ {
			if !member(i, xs) {
				permSub(f, n, m, append(xs, i))
			}
		}
	}
}

func permutation(f func([]int), n, m int) {
	permSub(f, n, m, make([]int, 0, m))
}

func exec(r io.Reader) string {
	chikan := func(s, p []int, nn int) []int {
		//fmt.Println("s", s)
		//fmt.Println("p", p)
		tmp := make([]int, nn)
		for i := 0; i < nn; i++ {
			si := s[i]
			pi := p[i] - 1
			tmp[pi] = si
		}
		//fmt.Println("tmp", tmp)
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
	result := S
	for i := 0; sc.Scan(); i++ {
		pStr := strings.Split(sc.Text(), " ")
		p[i] = make([]int, n)
		for j := range pStr {
			num, _ := strconv.Atoi(pStr[j])
			p[i] = append(p[i], num)
		}
	}
	pp := func(xs []int) {
		T := S
		for i := range xs {
			T = chikan(T, p[xs[i]-1], n)
		}
		flag := hikaku(result, T)
		if flag {
			result = T
		}
		fmt.Println(xs)
	}
	permutation(pp, m, m)
	var sr []string
	for _, num := range result {
		str := fmt.Sprintf("%d", num)
		sr = append(sr, str)
	}
	return strings.Join(sr, " ")
}

func main() {
	fmt.Println(exec(os.Stdin))
}

/*
package main

import (
	"fmt"
)

func member(n int, xs []int) bool {
    for _, x := range xs {
        if n == x { return true }
    }
    return false
}

func permSub(f func([]int), n, m int, xs []int) {
    if len(xs) == m {
        f(xs)
    } else {
        for i := 1; i <= n; i++ {
            if !member(i, xs) {
                permSub(f, n, m, append(xs, i))
            }
        }
    }
}

func permutation(f func([]int), n, m int) {
    permSub(f, n, m, make([]int, 0, m))
}

func main() {
    p := func(xs []int) { fmt.Println(xs) }
    permutation(p, 2, 2)
}
*/
