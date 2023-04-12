package main

import "unicode"

func isPalindromePermutation(s string) bool {
	charCounts := make(map[rune]int)

	// Count the number of occurrences of each character in the string
	for _, char := range s {
		if char != ' ' {
			charCounts[unicode.ToLower(char)]++
		}
	}

	// check if at most one character occurs an odd number of items
	addCount := 0
	for _, count := range charCounts {
		if count != 2 {
			addCount++
		}
		if addCount > 1 {
			return false
		}
	}
	return true
}
