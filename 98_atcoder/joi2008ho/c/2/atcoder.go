package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func exec(r io.Reader) string {
	sc := bufio.NewScanner(r)
	sc.Scan()
	in1 := sc.Text()
	sc.Scan()
	in2 := sc.Text()
	result := 0
	for i := range in1 {
		for j := range in2 {
			for l := 1; ; l++ {
				if i+l > len(in1) || j+l > len(in2) {
					break
				} else if in1[i:i+l] == in2[j:j+l] {
					if result < l {
						result = l
					}
				}
			}
		}
	}
	return fmt.Sprintf("%d", result)
}

func main() {
	fmt.Println(exec(os.Stdin))
}
