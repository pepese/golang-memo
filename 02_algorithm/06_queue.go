package main

import (
	"errors"
	"fmt"
)

type node struct {
	next *node
	data int
}

type queue struct {
	head, tail *node
	size       int
}

func (q *queue) len() int {
	return q.size
}

func (q *queue) enqueue(x int) {
	if q == nil || q.len() == 0 {
		q = newQueue(x)
		return
	}
	node := &node{data: x}
	if q.len() == 1 {
		q.head = node
		q.tail.next = q.head
	} else {
		q.head.next = node
		q.head = node
	}
	q.size++
}

func (q *queue) dequeue() (int, error) {
	if q == nil || q.len() == 0 {
		return 0, errors.New("error")
	}
	data := q.tail.data
	if q.len() == 1 {
		q.head, q.tail = nil, nil
	} else {
		q.tail = q.tail.next
	}
	q.size--
	return data, nil
}

func newQueue(x int) *queue {
	q := new(queue)
	node := &node{data: x}
	q.head, q.tail, q.size = node, node, 1
	return q
}

func main() {
	q := newQueue(1)
	q.enqueue(2)
	q.enqueue(3)
	num, err := q.dequeue()
	fmt.Println(num, err)
	num, err = q.dequeue()
	fmt.Println(num, err)
	num, err = q.dequeue()
	fmt.Println(num, err)
	num, err = q.dequeue()
	fmt.Println(num, err)

	var nilq queue
	fmt.Println("nilq", nilq)
	nilq.enqueue(1)
	fmt.Println("nilq", nilq)
}
