package main

import "fmt"

type Slice struct {
	ptr *[]int // Pointer to the underlying slice
	len int    // Length of the slice
	cap int    // Capacity of the slice
}

func NewSlice(cap int) *Slice {
	ptr := make([]int, cap)
	return &Slice{
		ptr: &ptr,
		len: 0,
		cap: cap,
	}
}

func (s *Slice) Append(val int) {
	if s.len == s.cap {
		newPtr := make([]int, s.cap*2)
		copy(newPtr, *s.ptr)
		fmt.Println(newPtr, *s.ptr)
		*s.ptr = append((*s.ptr)[:s.len], newPtr[s.len:]...)
		fmt.Println(newPtr, *s.ptr)
		s.cap *= 2
	}
	(*s.ptr)[s.len] = val
	s.len++
}

func (s *Slice) Print() {
	fmt.Printf("len =v%d cap = %d %v\n", s.len, s.cap, (*s.ptr)[:s.len])
}

func main() {
	// Create a new slice with a capacity of 3
	s := NewSlice(3)
	s.Print()

	// Append some values to the slice
	s.Append(1)
	s.Print()
	s.Append(2)
	s.Print()
	s.Append(3)
	s.Print()
	s.Append(4)
	s.Print()
}
