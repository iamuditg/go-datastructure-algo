# Generic Tree

This is an implementation of a generic tree data structure in Go. The program allows users to input a tree, perform various operations on it, and display the results. The supported operations include:

1. Insertion of the tree
2. Printing the tree in level order, pre-order, and post-order traversal
3. Computing the depth of the tree
4. Counting the number of nodes in the tree
5. Counting the number of leaf nodes in the tree
6. Deleting a node from the tree
7. Searching for a node in the tree

## Tree Traversals

- Level Order Traversal: This traversal visits the nodes of the tree level by level, starting from the root node and moving down to the deepest level. It prints the data of each node in level order.
- Pre-order Traversal: This traversal visits the nodes of the tree in the order of root, left child, and right child. It prints the data of each node in pre-order.
- Post-order Traversal: This traversal visits the nodes of the tree in the order of left child, right child, and root. It prints the data of each node in post-order.

## Function Descriptions

- `InsertTree()`: Prompts the user to enter the tree data, including the root node and its children, and returns the constructed tree.
- `LevelOrderTraversal(root *Node)`: Prints the tree in level order traversal.
- `PreOrderTraversal(root *Node)`: Prints the tree in pre-order traversal.
- `PostOrderTraversal(root *Node)`: Prints the tree in post-order traversal.
- `Depth(root *Node)`: Computes and returns the depth of the tree.
- `CountNodes(root *Node)`: Counts and returns the number of nodes in the tree.
- `CountLeafNodes(root *Node)`: Counts and returns the number of leaf nodes in the tree.
- `DeleteNode(root *Node, data string) bool`: Deletes the node with the specified data from the tree and returns true if the deletion is successful.
- `SearchNode(root *Node, data string) *Node`: Searches for a node with the specified data in the tree and returns the node if found, otherwise returns nil.

## Usage

1. Run the program.
2. Follow the instructions to input the tree data.
3. View the results of various operations on the tree.

Enjoy exploring the generic tree implementation!
