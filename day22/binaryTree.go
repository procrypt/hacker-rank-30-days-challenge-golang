package main

import (
	"errors"
	"fmt"
)

type Node struct {
	value int
	left  *Node
	right *Node
}

type Tree struct {
	root *Node
}

func(t *Tree) Insert(value int) error {
	if t.root == nil {
		t.root = &Node{value:value}
		return nil
	}
	return t.root.Insert(value)
}

func(t *Tree) Find(value int) (int, bool) {
	if t.root == nil {
		return 0, false
	}
	return t.root.Find(value)
}

func(n *Node) Insert(value int) error {
	if n == nil {
		return errors.New("nil tree")
	} else if  value == n.value {
		return nil
	} else if value < n.value {
		if n.left == nil {
			n.left = &Node{value:value}
			return nil
		} else {
			return n.left.Insert(value)
		}
	} else {
		if value > n.value {
			if n.right == nil {
				n.right = &Node{value:value}
				return nil
			} else {
				return n.right.Insert(value)
			}
		}
	}
	return nil
}

func(n *Node) Find(value int) (int, bool) {
	if n == nil {
		return 0, false
	} else if value == n.value {
		return n.value, true

	} else if value > n.value {
		return n.right.Find(value)

	} else {
		return n.left.Find(value)
	}
}


func main() {

	tree := &Tree{}

	tree.Insert(3)
	tree.Insert(2)
	tree.Insert(5)
	tree.Insert(1)
	tree.Insert(4)
	tree.Insert(6)

	fmt.Println(tree.root.value)
	fmt.Println(tree.root.left.value)
	fmt.Println(tree.root.right.value)
	fmt.Println(tree.root.left.left.value)
	fmt.Println(tree.root.right.left.value)
	fmt.Println(tree.root.right.right.value)
	fmt.Println(tree.Find(4))
}
