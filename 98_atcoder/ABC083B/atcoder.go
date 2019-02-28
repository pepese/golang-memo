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
	in := strings.Split(sc.Text(), " ")
	n, _ := strconv.Atoi(in[0])
	a, _ := strconv.Atoi(in[1])
	b, _ := strconv.Atoi(in[2])
	c := 0
	for i := 1; i <= n; i++ {
		if i == 10000 && a == 1 {
			c += i
		} else {
			tho := i / 1000
			hun := (i - tho*1000) / 100
			ten := (i - tho*1000 - hun*100) / 10
			one := i % 10
			if a <= tho+hun+ten+one && tho+hun+ten+one <= b {
				c += i
			}
		}
	}
	return fmt.Sprintf("%d", c)
}

func main() {
	fmt.Println(exec(os.Stdin))
}

/*
	if i >= 10000 {
		if a == 1 {
			c += 1
		}
	} else if i >= 1000 {
		tho := i / 1000
		hun := (i - tho*1000) / 100
		ten := (i - tho*1000 - hun*100) / 10
		one := i % 10
		if a <= tho+hun+ten+one && tho+hun+ten+one <= b {
			c += i
		}
	} else if i >= 100 {
		hun := i / 100
		ten := (i - hun*100) / 10
		one := i % 10
		if a <= hun+ten+one && hun+ten+one <= b {
			c += i
		}
	} else if i >= 10 {
		ten := i / 10
		one := i % 10
		if a <= ten+one && ten+one <= b {
			c += i
		}
	} else {
		if a <= i && i <= b {
			c += i
		}
	}*/
