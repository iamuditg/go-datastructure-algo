package main

import "fmt"

func main() {
	// Test case 1: Input is a permutation of a palindrome
	s1 := "Tact Coa"
	fmt.Println(isPalindromePermutation(s1)) // Output: true

	// Test case 2: Input is not a permutation of a palindrome
	s2 := "Hello Hello"
	fmt.Println(isPalindromePermutation(s2)) // Output: false
}
