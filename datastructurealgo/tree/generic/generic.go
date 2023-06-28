package main

import "fmt"

type Node struct {
	data     interface{}
	children []*Node
}

type Tree struct {
	root *Node
}

func NewTree() *Tree {
	return &Tree{}
}

func (t *Tree) Insert(data interface{}, parent *Node) {
	node := &Node{data, make([]*Node, 0)}
	if parent == nil {
		t.root = node
	} else {
		parent.children = append(parent.children, node)
	}
}

func (t *Tree) Print() {
	t.printHelper(t.root, 1)
}

func (t *Tree) printHelper(node *Node, level int) {
	if node == nil {
		return
	}
	fmt.Printf("level %d", level)
	for i, child := range node.children {
		fmt.Println("child-->", child.data)
		if i != len(node.children)-1 {
			fmt.Print(", ")
		}
	}
	fmt.Println()

	for _, child := range node.children {
		t.printHelper(child, level+1)
	}
}

func main() {
	tree := NewTree()

	// Inserting nodes level-wise
	tree.Insert(1, nil)
	tree.Insert(2, tree.root)
	tree.Insert(3, tree.root)
	tree.Insert(4, tree.root.children[0])
	tree.Insert(5, tree.root.children[0])
	tree.Insert(6, tree.root.children[1])

	// Printing all nodes
	tree.Print()
}
