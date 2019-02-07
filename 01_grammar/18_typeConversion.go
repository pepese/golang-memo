package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	snum := sc.Text()
	var num, strike, ball int
	num, _ = strconv.Atoi(snum)
	for i := 0; i < num; i++ {
		sc.Scan()
		w := sc.Text()
		if w == "strike" {
			strike++
			if strike < 3 {
				fmt.Println("strike!")
			} else {
				fmt.Println("out!")
			}
		} else if w == "ball" {
			ball++
			if ball < 4 {
				fmt.Println("ball!")
			} else {
				fmt.Println("fourball!")
			}
		}
	}
}
