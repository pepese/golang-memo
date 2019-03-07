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
	n, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	in := strings.Split(sc.Text(), " ")
	arr := make([]int, n)
	for i, ai := range in {
		tmp, _ := strconv.Atoi(ai)
		arr[i] = tmp
	}
	sort.Sort(sort.IntSlice(arr))
	result := make([]int, n/2)
	for i := range result {
		result[i] = arr[i] + arr[len(arr)-1-i]
	}
	sort.Sort(sort.IntSlice(result))
	return fmt.Sprintf("%d", result[len(result)-1]-result[0])
}

func main() {
	fmt.Println(exec(os.Stdin))
}
