package main

import "fmt"

type Node struct {
	value       int
	left, right *Node
}

type BST struct {
	root *Node
}

func NewBST() *BST {
	return &BST{}
}

func (bst *BST) Insert(value int) {
	node := &Node{value: value}
	if bst.root == nil {
		bst.root = node
	} else {
		insertNode(bst.root, node)
	}
}

func insertNode(root *Node, node *Node) {
	if node.value < root.value {
		if root.left == nil {
			root.left = node
		} else {
			insertNode(root.left, node)
		}
	} else {
		if root.right == nil {
			root.right = node
		} else {
			insertNode(root.right, node)
		}
	}
}

func searchNode(root *Node, value int) bool {
	if root == nil {
		return false
	}
	if value == root.value {
		return true
	} else if value < root.value {
		return searchNode(root.left, value)
	} else {
		return searchNode(root.right, value)
	}
}

func (bst *BST) Delete(value int) {
	bst.root = deleteNode(bst.root, value)
}

func deleteNode(root *Node, value int) *Node {
	if root == nil {
		return nil
	}

	if value < root.value {
		root.left = deleteNode(root.left, value)
	} else if value > root.value {
		root.right = deleteNode(root.right, value)
	} else {
		if root.left == nil && root.right == nil {
			root = nil
		} else if root.left == nil {
			root = root.right
		} else if root.right == nil {
			root = root.left
		} else {
			minRight := findMinValue(root.right)
			root.value = minRight
			root.right = deleteNode(root.right, minRight)
		}
	}
	return root
}

func findMinValue(root *Node) int {
	minVal := root.value
	for root.left != nil {
		minVal = root.left.value
		root = root.left
	}
	return minVal
}

func (bst *BST) InOrderTraversal() {
	inOrder(bst.root)
	fmt.Println()
}

func inOrder(node *Node) {
	if node != nil {
		inOrder(node.left)
		fmt.Printf("%d ", node.value)
		inOrder(node.right)
	}
}

func (bst *BST) Search(value int) bool {
	return searchNode(bst.root, value)
}

func main() {
	bst := NewBST()

	bst.Insert(50)
	bst.Insert(30)
	bst.Insert(70)
	bst.Insert(20)
	bst.Insert(40)
	bst.Insert(60)
	bst.Insert(80)

	bst.InOrderTraversal()

	fmt.Println("Is 40 present?", bst.Search(40))
	fmt.Println("Is 90 present?", bst.Search(90))

	bst.Delete(30)
	bst.Delete(80)

	bst.InOrderTraversal()
}
