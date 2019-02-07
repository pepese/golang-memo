package main

// C060:辞書の作成

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	s := sc.Text()
	nums := strings.Split(s, " ")
	//num1, _ := strconv.Atoi(nums[0])
	num2, _ := strconv.Atoi(nums[1])
	num3, _ := strconv.Atoi(nums[2])
	//nn := num1 % num2
	sc.Scan()
	s = sc.Text()
	words := strings.Split(s, " ")
	sort.Strings(words)
	for i, word := range words {
		if i >= num2*(num3-1) && i < num2*num3 {
			fmt.Println(word)
		}
	}
}
