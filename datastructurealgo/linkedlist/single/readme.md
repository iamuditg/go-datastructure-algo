# Singly Linked List

A singly linked list is a data structure that consists of a sequence of nodes. Each node contains data and a reference (link) to the next node in the sequence.

## Implementation Details

The provided Go code implements a singly linked list with the following functionalities:

1. `InsertElement(value int)` - Inserts an element at the end of the linked list.
2. `InsertAt(value int, index int)` - Inserts an element at the specified index in the linked list.
3. `DeleteAt(index int)` - Deletes the element at the specified index from the linked list.
4. `DeleteAll()` - Deletes all elements from the linked list.
5. `ReverseLinkedList()` - Reverses the order of elements in the linked list.
6. `MergeSorted(other *LinkedList)` - Merges the current linked list with another linked list in sorted order.
7. `MergeSort()` - Sorts the linked list in ascending order using the merge sort algorithm.
8. `SearchElement(value int)` - Searches for an element in the linked list and returns true if found, false otherwise.
9. `MiddleElement()` - Returns the middle element of the linked list.

## Usage

To use the provided implementation, follow these steps:

1. Create a new instance of the linked list: `ll := NewLinkedList()`
2. Insert elements using `InsertElement()` or `InsertAt()` methods.
3. Perform operations such as deletion, reversal, merging, and searching as required.
4. Access the middle element using `MiddleElement()`.
5. Delete all elements using `DeleteAll()` when no longer needed.

Please refer to the code comments for more details on each function and its usage.

## Example

Here's an example usage of the provided singly linked list implementation:

```go
package main

import (
	"fmt"
)

func main() {
	// Create a new linked list
	ll := NewLinkedList()

	// Insert elements into the linked list
	ll.InsertElement(5)
	ll.InsertElement(10)
	ll.InsertElement(3)

	// Insert an element at a specific index
	ll.InsertAt(7, 1)

	// Delete an element at a specific index
	ll.DeleteAt(2)

	// Reverse the linked list
	ll.ReverseLinkedList()

	// Merge two linked lists in sorted order
	otherList := NewLinkedList()
	otherList.InsertElement(2)
	otherList.InsertElement(8)
	ll.MergeSorted(otherList)

	// Search for an element in the linked list
	fmt.Println("Search Result:", ll.SearchElement(4)) // Output: true
	fmt.Println("Search Result:", ll.SearchElement(7)) // Output: false

	// Get the middle element of the linked list
	middle := ll.MiddleElement()
	fmt.Println("Middle Element:", middle.value) // Output: 10
}
