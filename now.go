package main

import (
	"fmt"
	"sort"
)

type user struct {
	name string
	age  int
}

type users []user

func (u users) Len() int {
	return len(u)
}

func (u users) Less(i, j int) bool {
	return u[i].age < u[j].age
}

func (u users) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func main() {
	list := []user{
		user{name: "A", age: 2},
		user{name: "B", age: 1},
		user{name: "C", age: 4},
		user{name: "D", age: 3},
	}
	us := users(list)
	sort.Sort(us)
	fmt.Println(us)
}
