package main

import "fmt"

type Node struct {
	val int
	left, right *Node
}

func (node *Node) insert(val int) *Node {
	if node == nil {
		return &Node{val: val} // Initial root node
	}

	if val < node.val {
		node.left = node.left.insert(val)
	} else if val > node.val {
		node.right = node.right.insert(val)
	}
	return node
}

func (node *Node) contains(val int) bool {
	switch {
		case node == nil:
			return false
		case val < node.val:
			return node.left.contains(val)
		case val > node.val:
			return node.right.contains(val)
		default:
			return true
	}
}

func main() {
	var root *Node
	root = root.insert(5)
	root = root.insert(3)
	root = root.insert(10)
	root = root.insert(2)
	fmt.Println(root.contains(2))
	fmt.Println(root.contains(12))
}
