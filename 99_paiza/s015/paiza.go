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
func exec(r io.Reader) string {
	sc := bufio.NewScanner(r)
	sc.Scan()
	ins := strings.Split(sc.Text(), " ")
	k, _ := strconv.Atoi(ins[0])
	s, _ := strconv.Atoi(ins[1])
	t, _ := strconv.Atoi(ins[2])
	level := "ABC"
	for i := 1; i < k; i++ {
		level = "A" + level + "B" + level + "C"
		fmt.Println(level)
	}
	return level[s-1 : t]
}
*/
/*
func exec(r io.Reader) string {
	sc := bufio.NewScanner(r)
	sc.Scan()
	ins := strings.Split(sc.Text(), " ")
	k, _ := strconv.Atoi(ins[0])
	s, _ := strconv.Atoi(ins[1])
	t, _ := strconv.Atoi(ins[2])
	l := []byte{65, 66, 67}
	for i := 1; i < k; i++ {
		a := append([]byte{65}, l...)
		b := append([]byte{66}, l...)
		b = append(b, []byte{67}...)
		l = append(a, b...)
	}
	return string(l[s-1 : t])
}
*/
/*
func exec(r io.Reader) string {
	en := func(n int) int {
		return 3 * (int(math.Pow(2.0, float64(n+1))) - 1)
	}
	sc := bufio.NewScanner(r)
	sc.Scan()
	ins := strings.Split(sc.Text(), " ")
	k, _ := strconv.Atoi(ins[0])
	s, _ := strconv.Atoi(ins[1])
	t, _ := strconv.Atoi(ins[2])
	l := make([]byte, en(k-1), 2*en(k-1))
	top := k - 1
	l[top], l[top+1], l[top+2] = 65, 66, 67
	for i := 0; i < k-1; i++ {
		size := en(i)
		l[top-1] = 65
		l[top+size] = 66
		for j := top + size + 1; j < top+2*size+1; j++ {
			l[j] = l[j-size-1]
		}
		l[top+2*size+1] = 67
		top--
	}
	return string(l[s-1 : t])
}
*/
/*
func exec(r io.Reader) string {
	sc := bufio.NewScanner(r)
	sc.Scan()
	ins := strings.Split(sc.Text(), " ")
	k, _ := strconv.Atoi(ins[0])
	s, _ := strconv.Atoi(ins[1])
	t, _ := strconv.Atoi(ins[2])
	l := []byte{65, 66, 67}
	for i := 0; i < k-1; i++ {
		var cp []byte
		cp = append(cp, l...)
		// A 追加
		a := []byte{65}
		l = append(a, l...)
		// B 、前レベル追加
		l = append(l, 66)
		l = append(l, cp...)
		// C 追加
		l = append(l, 67)
		//fmt.Println(l)
	}
	return string(l[s-1 : t])
}
*/
/*
func exec(r io.Reader) string {
	en := func(n int) int {
		return 3 * (int(math.Pow(2.0, float64(n+1))) - 1)
	}
	sc := bufio.NewScanner(r)
	sc.Scan()
	ins := strings.Split(sc.Text(), " ")
	k, _ := strconv.Atoi(ins[0])
	s, _ := strconv.Atoi(ins[1])
	t, _ := strconv.Atoi(ins[2])
	l := make([]*byte, en(k-1))
	var a, b, c byte
	a, b, c = 65, 66, 67
	top := k - 1
	l[top], l[top+1], l[top+2] = &a, &b, &c
	for i := 0; i < k-1; i++ {
		size := en(i)
		l[top-1] = &a
		l[top+size] = &b
		for j := top + size + 1; j < top+2*size+1; j++ {
			l[j] = l[j-size-1]
		}
		l[top+2*size+1] = &c
		top--
	}
	var result []byte
	for i := s - 1; i < t; i++ {
		result = append(result, *l[i])
	}
	//return string(l[s-1 : t])
	return string(result)
}
*/
func exec(r io.Reader) string {
	sc := bufio.NewScanner(r)
	sc.Scan()
	ins := strings.Split(sc.Text(), " ")
	k, _ := strconv.Atoi(ins[0])
	s, _ := strconv.Atoi(ins[1])
	t, _ := strconv.Atoi(ins[2])
	var a, b, c byte
	a, b, c = 65, 66, 67
	l := []*byte{&a, &b, &c}
	for i := 0; i < k-1; i++ {
		var cp []*byte
		cp = append(cp, l...)
		// A 追加
		a := []*byte{&a}
		l = append(a, l...)
		// B 、前レベル追加
		l = append(l, &b)
		l = append(l, cp...)
		// C 追加
		l = append(l, &c)
		//fmt.Println(l)
	}
	var result []byte
	for i := s - 1; i < t; i++ {
		result = append(result, *l[i])
	}
	//return string(l[s-1 : t])
	return string(result)
}

func main() {
	fmt.Println(exec(os.Stdin))
}
