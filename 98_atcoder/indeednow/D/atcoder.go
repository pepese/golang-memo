package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type node struct {
	from, to []*node
}

func exec(r io.Reader) string {
	sc := bufio.NewScanner(r)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())
	nodes := make([]*node, n+1)
	for i := 0; i < n-1; i++ {
		sc.Scan()
		in := strings.Split(sc.Text(), " ")
		a, _ := strconv.Atoi(in[0])
		b, _ := strconv.Atoi(in[1])
		t, _ := strconv.Atoi(in[2])
		if nodes[a] == nil {
			nodes[a] = &node{}
		}
		if nodes[b] == nil {
			nodes[b] = &node{}
		}
		if nodes[a].to == nil {
			nodes[a].to = []*node{}
		}
		nodes[a].to = append(nodes[a].to, nodes[b])
		if nodes[b].from == nil {
			nodes[b].from = []*node{}
		}
		nodes[b].from = append(nodes[b].from, nodes[a])
		if t == 2 {
			if nodes[a].from == nil {
				nodes[a].from = []*node{}
			}
			nodes[a].from = append(nodes[a].from, nodes[b])
			if nodes[b].to == nil {
				nodes[b].to = []*node{}
			}
			nodes[b].to = append(nodes[b].to, nodes[a])
		}
	}
	result := 0
	gmemo := make(map[string]int, n)
	var search func(from, nd *node, c int, memo map[*node]int)
	search = func(from, nd *node, c int, memo map[*node]int) {
		if nd == nil {
			return
		}
		if nd.to == nil && result < c-1 {
			result = c - 1

		}
		if nd.to != nil && len(nd.to) > 0 {
			for i, to := range nd.to {
				dik := gmemo[fmt.Sprintf("%v", to)]
				if memo[to] != 1 && dik <= c {
					gmemo[fmt.Sprintf("%v", to)] = c
					memo[to] = 1
					search(from, nd.to[i], c+1, memo)
				}
			}
		}
	}
	for i := 1; i < len(nodes); i++ {
		mp := make(map[*node]int)
		mp[nodes[i]] = 1
		search(nodes[i], nodes[i], 0, mp)
	}
	return fmt.Sprintf("%d", result)
}

func main() {
	fmt.Println(exec(os.Stdin))
}
