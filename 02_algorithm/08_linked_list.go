package main

type node struct {
	next *node
	data int
}

type linkedList struct {
	root *node
	size int
}

func (l *linkedList) len() int {
	return l.size
}

func (l *linkedList) append(x int) {
	node := &node{data: x}
	if l.len() == 0 {
		l.root, l.size = node, 1
	} else {
		n := l.root
		for n.next != nil {
			n = n.next
		}
		n.next = node
		l.size++
	}
}

func (l *linkedList) remove(x int) bool {
	if l.len() == 0 {
		return false
	}
	n := l.root
	for n != nil {
		if n.data == x {
			return true
		}
		n = n.next
	}
	return false
}
