package main

import "fmt"

type Node struct {
	value int
	left  *Node
	right *Node
}

func newNode(x int) *Node {
	node := new(Node)
	node.value = x
	return node
}

func insertNode(node *Node, x int) *Node {
	if node == nil {
		return newNode(x)
	}
	if node.value > x {
		node.left = insertNode(node.left, x)
	} else {
		node.right = insertNode(node.right, x)
	}
	return node
}

func searchNode(node *Node, x int) bool {
	n := node
	for n != nil {
		switch {
		case n.value == x:
			return true
		case n.value > x:
			n = n.left
		default:
			n = n.right
		}
	}
	return false
}

func main() {
	root := newNode(5)
	fmt.Println(*root)
	root = insertNode(root, 4)
	fmt.Println(root.left)
	fmt.Println(root.right)
	fmt.Println(searchNode(root, 4))
}
