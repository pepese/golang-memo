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

type person struct {
	number, vote int
}

func exec(r io.Reader) string {
	sc := bufio.NewScanner(r)
	sc.Scan()
	ins := strings.Split(sc.Text(), " ")
	m, _ := strconv.Atoi(ins[0]) // 立候補者
	n, _ := strconv.Atoi(ins[1]) // 有権者
	//k, _ := strconv.Atoi(ins[2]) // 演説回数
	people := make([]*person, m)
	for i := 0; i < m; i++ {
		people[i] = &person{
			number: i + 1,
			vote:   0,
		}
	}
	for sc.Scan() {
		e, _ := strconv.Atoi(sc.Text())
		eNumber := e - 1
		move := 0
		for i, p := range people {
			if p.vote > 0 && i != eNumber {
				move++
				p.vote--
			}
		}
		if n > 0 {
			people[eNumber].vote += move + 1
			n--
		} else {
			people[eNumber].vote += move
		}
	}
	sort.SliceStable(people, func(i, j int) bool {
		return people[i].number < people[j].number
	})
	sort.SliceStable(people, func(i, j int) bool {
		return people[i].vote > people[j].vote
	})
	max := people[0].vote
	var tops []string
	for _, p := range people {
		if p.vote == max {
			tops = append(tops, strconv.Itoa(p.number))
		} else {
			break
		}
	}

	return strings.Join(tops, "\n")
}

func main() {
	fmt.Println(exec(os.Stdin))
}
