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

func main() {
	fmt.Println(exec(os.Stdin))
}
