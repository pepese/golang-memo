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
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	a := strings.Split(sc.Text(), " ")
	an := make([]int, n)
	for i, m := range a {
		an[i], _ = strconv.Atoi(m)
	}
	for i := 0; true; i++ {
		tmp := make([]int, 0, n)
		for j := 0; j < n; j++ {
			if an[j]%2 == 0 {
				tmp = append(tmp, an[j]/2)
			} else {
				break
			}
		}
		if len(tmp) < n {
			return fmt.Sprintf("%d", i)
		} else {
			an = tmp
		}
	}
	return ""
}

func main() {
	fmt.Println(exec(os.Stdin))
}
