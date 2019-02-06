package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var in int
	fmt.Fscanf(os.Stdin, "%d\n", &in)
	var s []string
	for i := 0; i < in; i++ {
		s = append(s, "*")
	}
	fmt.Println(strings.Join(s, ""))
}
