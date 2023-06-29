package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Data     string
	Children []*Node
}

type Tree struct {
	Root *Node
}

func InsertTree() *Tree {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the root node:")
	rootData, _ := reader.ReadString('\n')
	rootData = strings.TrimSpace(rootData)

	root := &Node{Data: rootData, Children: []*Node{}}
	tree := &Tree{Root: root}

	queue := []*Node{root}
	fmt.Println("queue[0]", queue[0], queue[1:])
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		fmt.Printf("Enter the number of children for node %s (or 0 to skip): ", node.Data)
		childCountStr, _ := reader.ReadString('\n')
		childCountStr = strings.TrimSpace(childCountStr)
		childCount, _ := strconv.Atoi(childCountStr)

		if childCount == 0 {
			continue
		}

		for i := 0; i < childCount; i++ {
			fmt.Printf("Enter child %d:", i+1)
			childData, _ := reader.ReadString('\n')
			childData = strings.TrimSpace(childData)
			InsertChildren(node, []string{childData})
		}
		queue = append(queue, node.Children...)
	}
	return tree
}

func InsertChildren(parent *Node, childData []string) {
	for _, data := range childData {
		child := &Node{
			Data:     data,
			Children: []*Node{},
		}
		parent.Children = append(parent.Children, child)
	}
}

func PostOrderTraversal(root *Node) {
	if root == nil {
		return
	}

	for _, child := range root.Children {
		PostOrderTraversal(child)
	}
	fmt.Printf("%s", root.Data)
}

func PreOrderTraversal(root *Node) {
	if root == nil {
		return
	}

	fmt.Printf("%s ", root.Data)

	for _, child := range root.Children {
		PreOrderTraversal(child)
	}
}

func LevelOrderTraversal(root *Node) {
	if root == nil {
		return
	}

	queue := []*Node{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		fmt.Println("len:", levelSize)
		for i := 0; i < levelSize; i++ {
			node := queue[i]
			fmt.Printf("%s ", node.Data)
			for _, child := range node.Children {
				queue = append(queue, child)
			}
		}
		fmt.Println()
		queue = queue[levelSize:]
	}
}

func Depth(root *Node) int {
	if root == nil {
		return 0
	}

	maxDepth := 0
	for _, child := range root.Children {
		childDepth := Depth(child)
		if childDepth > maxDepth {
			maxDepth = childDepth
		}
	}
	return maxDepth + 1
}

func CountNodes(root *Node) int {
	if root == nil {
		return 0
	}
	count := 1
	for _, child := range root.Children {
		count += CountNodes(child)
	}
	return count
}

func CountLeafNodes(root *Node) int {
	if root == nil {
		return 0
	}
	if len(root.Children) == 0 {
		return 1
	}
	count := 0
	for _, child := range root.Children {
		count += CountLeafNodes(child)
	}
	return count
}

func DeleteNode(root *Node, data string) bool {
	if root == nil {
		return false
	}

	if root.Data == data {
		root = nil
		return true
	}

	for i, child := range root.Children {
		if child.Data == data {
			root.Children = append(root.Children[:i], root.Children[i+1:]...)
			return true
		}

		if DeleteNode(child, data) {
			return true
		}
	}

	return false
}

func SearchNode(root *Node, data string) *Node {
	if root == nil {
		return nil
	}

	if root.Data == data {
		return root
	}

	for _, child := range root.Children {
		foundNode := SearchNode(child, data)
		if foundNode != nil {
			return foundNode
		}
	}

	return nil
}

func main() {
	fmt.Println("=== Generic Tree Implementation ===")
	fmt.Println("Enter the tree:")
	tree := InsertTree()
	fmt.Println(tree, tree.Root, tree.Root.Data, tree.Root.Children)

	fmt.Println("\n Post Order Traversal")
	PostOrderTraversal(tree.Root)
	println()

	fmt.Println("\nLevel Order Traversal:")
	LevelOrderTraversal(tree.Root)
	fmt.Println()

	fmt.Println("\nPre Order Traversal:")
	PreOrderTraversal(tree.Root)
	fmt.Println()

	fmt.Println("\nPre Order Traversal:")
	PreOrderTraversal(tree.Root)
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nEnter a node data to delete:")
	nodeData, _ := reader.ReadString('\n')
	nodeData = strings.TrimSpace(nodeData)
	if DeleteNode(tree.Root, nodeData) {
		fmt.Println("Node deleted successfully.")
		fmt.Println("\nLevel Order Traversal after deletion:")
		LevelOrderTraversal(tree.Root)
		fmt.Println()
	} else {
		fmt.Println("Node not found.")
	}

	fmt.Println("\nEnter a node data to search:")
	nodeData, _ = reader.ReadString('\n')
	nodeData = strings.TrimSpace(nodeData)
	if node := SearchNode(tree.Root, nodeData); node != nil {
		fmt.Printf("Node found: %s\n", node.Data)
	} else {
		fmt.Println("Node not found.")
	}

}
