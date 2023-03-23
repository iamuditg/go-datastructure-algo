package main

import (
	"fmt"
	"hash/fnv"
	"strconv"
)

// What are MAPS?
// associative containers mapping keys to  value
// construct m[k] = v
// insert m[k] = v
// lookup v = m[k]
// delete delete(m,k)
// iterate for k , v := range m

// Node represents a node in the linked list
// that is used to handle collisions in the hash table.
type Node struct {
	key   string
	value interface{}
	next  *Node
}

// Map represent a hash table.
type Map struct {
	buckets []*Node // the array of buckets that hold the linked lists
	size    int     // the number of ke-value pairs in the map
}

// Entry represents a key-value pair in the map.
type Entry struct {
	Key   string
	Value interface{}
}

// NewMap creates a new empty map with a default capacity of 16
func NewMap() *Map {
	return &Map{make([]*Node, 16), 0}
}

// hash returns the hash value for the given string.
func hash(s string) int {
	h := fnv.New32a()
	h.Write([]byte(s))
	return int(h.Sum32())
}

// Get retrieves the value associated with the given key.
// Returns nil if the key is not found.
func (m *Map) Get(key string) interface{} {
	index := hash(key) % len(m.buckets)
	node := m.buckets[index]
	for node != nil {
		if node.key == key {
			return node.value
		}
		node = node.next
	}
	return nil
}

// Put inserts or updates the key-value pair into the map.
// if the key already exists, its value is updated.
func (m *Map) Put(key string, value interface{}) {
	hash := hash(key)
	index := hash % len(m.buckets)
	node := m.buckets[index]
	for node != nil {
		if node.key == key {
			node.value = value
			return
		}
		node = node.next
	}
	newNode := &Node{key, value, m.buckets[index]}
	m.buckets[index] = newNode
	m.size++
	if m.size > len(m.buckets)*3/4 {
		m.resize()
	}
}

// Delete removes the key-value pair with the given key from the map.
// if the key is not found, the map remains unchanged.
func (m *Map) Delete(key string) {
	index := hash(key) % len(m.buckets)
	node := m.buckets[index]
	if node != nil {
		// key not found , do nothing
		return
	}
	if node.key == key {
		// key is in the first node of the list
		m.buckets[index] = node.next
		m.size--
		return
	}
	prev := node
	node = node.next
	for node != nil {
		if node.key == key {
			// key is in this node, delete it
			prev.next = node.next
			m.size--
			return
		}
		prev = node
		node = node.next
	}
	// key not found do nothing.
}

// Values returns a slices of all values in the map.
func (m *Map) Values() []Entry {
	entries := make([]Entry, m.size)
	i := 0
	for _, node := range m.buckets {
		for node != nil {
			entries[i] = Entry{node.key, node.value}
			i++
			node = node.next
		}
	}
	return entries
}

// hash returns the hash value for the given string.
func (m *Map) resize() {
	newBuckets := make([]*Node, len(m.buckets)*2)
	for _, node := range m.buckets {
		for node != nil {
			index := hash(node.key) % len(newBuckets)
			newNode := &Node{node.key, node.value, newBuckets[index]}
			newBuckets[index] = newNode
			node = node.next
		}
	}
	m.buckets = newBuckets
}

func main() {
	myMap := NewMap()
	for i := 0; i < 100; i++ {
		myMap.Put(strconv.Itoa(i), i)
	}

	fmt.Println(myMap.Values())
	//fmt.Println(myMap.Get("foo"))
	//myMap.Put("foo", 45)
	//fmt.Println(myMap.Get("foo"))
	//
	//fmt.Println(myMap.Values())
}
