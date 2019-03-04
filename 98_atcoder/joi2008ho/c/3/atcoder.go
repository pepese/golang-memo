package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

func exec(r io.Reader) string {
	sc := bufio.NewScanner(r)
	sc.Scan()
	in := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(in[0])
	m, _ := strconv.Atoi(in[1])
	pi := make([]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		p, _ := strconv.Atoi(sc.Text())
		pi[i] = p
	}
	sort.Sort(sort.Reverse(sort.IntSlice(pi)))
	result := 0
	for _, n1 := range pi {
		if n1 > m {
			break
		} else if result < n1 {
			result = n1
		} else if n1+3*pi[0] < result {
			break
		}
		for _, n2 := range pi {
			if n1+n2 > m {
				break
			} else if result < n1+n2 {
				result = n1 + n2
			} else if n1+n2+2*pi[0] < result {
				break
			}
			for _, n3 := range pi {
				if n1+n2+n3 > m {
					break
				} else if result < n1+n2+n3 {
					result = n1 + n2 + n3
				} else if n1+n2+n3+pi[0] < result {
					break
				}
				for _, n4 := range pi {
					if n1+n2+n3+n4 > m {
						break
					} else if result < n1+n2+n3+n4 {
						result = n1 + n2 + n3 + n4
						break
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
