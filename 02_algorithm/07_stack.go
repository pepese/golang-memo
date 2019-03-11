package main

import (
	"errors"
	"fmt"
)

type node struct {
	next *node
	data int
}

type stack struct {
	top  *node
	size int
}

func (s *stack) len() int {
	return s.size
}

func (s *stack) push(x int) {
	node := &node{data: x}
	if s.len() == 0 {
		s.top, s.size = node, 1
	} else {
		node.next, s.top = s.top, node
		s.size++
	}
}

func (s *stack) pop() (int, error) {
	if s == nil || s.len() == 0 {
		return 0, errors.New("error")
	}
	data := s.top.data
	s.top = s.top.next
	s.size--
	return data, nil
}

func newStack(x int) *stack {
	s := new(stack)
	node := &node{data: x}
	s.top, s.size = node, 1
	return s
}

func main() {
	s := newStack(1)
	s.push(2)
	s.push(3)
	num, err := s.pop()
	fmt.Println(num, err)
	num, err = s.pop()
	fmt.Println(num, err)
	num, err = s.pop()
	fmt.Println(num, err)
	num, err = s.pop()
	fmt.Println(num, err, s.top)
	s.push(4)
	num, err = s.pop()
	fmt.Println(num, err)

	var nils stack
	fmt.Println("nils", nils)
	num, err = nils.pop()
	fmt.Println(num, err)
	nils.push(10)
	fmt.Println("nils", nils)
	num, err = nils.pop()
	fmt.Println(num, err)
}
