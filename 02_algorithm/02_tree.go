package main

import (
	"fmt"
)

type binaryNode struct {
	left, right *binaryNode
	data        int
}

func (n *binaryNode) insert(x int) {
	if n == nil {
		return
	}
	if n.data > x {
		if n.left == nil {
			n.left = newBinaryNode(x)
		} else {
			n.left.insert(x)
		}
	} else {
		if n.right == nil {
			n.right = newBinaryNode(x)
		} else {
			n.right.insert(x)
		}
	}
}

func (n *binaryNode) search(x int) bool {
	for n != nil {
		switch {
		case n.data == x:
			return true
		case n.data > x:
			n = n.left
		default:
			n = n.right
		}
	}
	return false
}

func newBinaryNode(x int) *binaryNode {
	node := new(binaryNode)
	node.data = x
	return node
}

func main() {
	root := newBinaryNode(5)
	fmt.Println(*root)
	root.insert(4)
	fmt.Println(root.left)
	fmt.Println(root.right)
	fmt.Println(root.search(4))
}
