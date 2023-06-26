# Generic Tree Implementation

This is an implementation of a generic tree data structure in Go. The tree supports various operations such as insertion, printing, deletion, calculating height, printing all nodes, and counting the number of nodes.

## Implementation Details

The generic tree is represented using the following components:

- **Node:** The `Node` struct represents a node in the tree. It contains a value and a slice of child nodes.

- **Tree:** The `Tree` struct represents the generic tree. It contains a root node and provides methods for tree operations.

The following operations are implemented in the tree:

1. **Insert:** The `Insert` method is used to insert a value into the tree. It takes a value and an optional parent node. If the parent node is `nil`, the value is inserted as the root of the tree. Otherwise, it is added as a child node to the parent.

2. **Print:** The tree supports four print methods: `PreOrder`, `InOrder`, `PostOrder`, and `LevelOrder`. These methods traverse the tree in different orders and print the values of the nodes.

3. **Delete:** The `Delete` method is used to delete a node from the tree. It takes a value as input and deletes the node with the matching value from the tree. If the node has children, they are also deleted.

4. **Height:** The `Height` method calculates the height of the tree. It returns the maximum number of levels in the tree.

5. **Print All Nodes:** The `PrintAllNodes` method prints all the nodes in the tree. It uses in-order traversal to visit the nodes.

6. **Count:** The `Count` method returns the total number of nodes in the tree.

## Usage

To use the generic tree implementation, follow these steps:

1. Create a new tree using the `NewTree` function.

2. Insert nodes into the tree using the `Insert` method.

3. Perform tree operations such as printing, deletion, calculating height, printing all nodes, or counting the number of nodes.

Here's an example of using the generic tree implementation:

```go
tree := NewTree()

// Insert nodes into the tree
tree.Insert("A", nil)
tree.Insert("B", tree.root)
tree.Insert("C", tree.root)
tree.Insert("D", tree.root.children[0])
tree.Insert("E", tree.root.children[0])
tree.Insert("F", tree.root.children[1])
tree.Insert("G", tree.root.children[2])

// Print the tree using different traversal methods
tree.PreOrder(tree.root)
tree.InOrder(tree.root)
tree.PostOrder(tree.root)
tree.LevelOrder()

// Calculate the height of the tree
tree.Height()

// Print all nodes in the tree
tree.PrintAllNodes()

// Count the total number of nodes in the tree
tree.Count()

// Delete a node from the tree
tree.Delete("D")
