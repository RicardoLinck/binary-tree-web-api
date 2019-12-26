package binarytree

import (
	"fmt"
	"strconv"
	"strings"
)

// Node represents a binary tree node
type Node struct {
	Value int
	Left  *Node
	Right *Node
}

// BinaryTree represents the main pointers to Nodes
type BinaryTree struct {
	Root *Node
}

// Contains function checks if the value is present in the tree
func (bt *BinaryTree) Contains(value int) bool {
	if bt.Root == nil {
		return false
	}
	return bt.Root.contains(value)
}

func (n *Node) contains(value int) bool {
	if value < n.Value {
		if n.Left == nil {
			return false
		}
		return n.Left.contains(value)

	} else if value == n.Value {
		return true
	} else {
		if n.Right == nil {
			return false
		}
		return n.Right.contains(value)
	}
}

// Insert the value to the root, if the root is null, creates the root with value
func (bt *BinaryTree) Insert(value int) {
	if bt.Root == nil {
		bt.Root = newNode(value)
	} else {
		bt.Root.insert(value)
	}
}

// Insert a value to the left node in case the value is lesser than node value
// or to the right node otherwise
func (n *Node) insert(value int) {
	if value < n.Value {
		if n.Left == nil {
			n.Left = newNode(value)
		} else {
			n.Left.insert(value)
		}
	} else if value == n.Value {
		return
	} else {
		if n.Right == nil {
			n.Right = newNode(value)
		} else {
			n.Right.insert(value)
		}
	}
}

// PrintLtr prints the elements of the binary tree from left to right
func (bt *BinaryTree) PrintLtr() string {
	return bt.Root.printLtr()
}

func (n *Node) printLtr() string {
	if n == nil {
		return ""
	}
	return strings.Trim(fmt.Sprintf("%s %s %s", n.Left.printLtr(), strconv.Itoa(n.Value), n.Right.printLtr()), " ")
}

func newNode(value int) *Node {
	return &Node{Value: value}
}
