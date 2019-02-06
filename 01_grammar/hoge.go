package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	var in string
	fmt.Fscanf(os.Stdin, "%s", &in)
	fmt.Println(strings.Replace(in, " ", "/", 2))
	/*
		sc := bufio.NewScanner(os.Stdin)
		sc.Split(bufio.ScanWords)
		sc.Scan()
		a := sc.Text()
		sc.Scan()
		b := sc.Text()
		sc.Scan()
		c := sc.Text()
		fmt.Printf("%s/%s/%s\n", a, b, c)
	*/
}
