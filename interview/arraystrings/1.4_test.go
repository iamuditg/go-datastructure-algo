package main

import (
	"testing"
	"unicode"
)

/*
Palindrome Permutation: Given a string, write a function to check if it is a permutation of a palindrome. A palindrome is a word or phrase that is the same forwards and backwards. A permutation
is a rearrangement of letters. The palindrome does not need to be limited to just dictionary words.
1.5
1.6
EXAMPLE
Input: Tact Coa
Output: True (permutations: "taco cat", "atco eta", etc.)
*/

func isPalindromePermutation(s string) bool {
	charCounts := make(map[rune]int)

	// Count the number of occurrences of each character in the string
	for _, char := range s {
		if char != ' ' {
			charCounts[unicode.ToLower(char)]++
		}
	}

	// Check if at most one character occurs an odd number of times
	addCount := 0
	for _, count := range charCounts {
		if count%2 != 0 {
			addCount++
		}
		if addCount > 1 {
			return false
		}
	}

	return true
}

func TestIsPalindromePermutation(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"Tact Coa", true},
		{"hello", false},
		{"aab", true},
		{"abcd", false},
	}

	for _, test := range tests {
		result := isPalindromePermutation(test.input)
		if result != test.expected {
			t.Errorf("isPalindromePermutation(%q) = %v, expected %v", test.input, result, test.expected)
		}
	}
}

func BenchmarkIsPalindromePermutation(b *testing.B) {
	input := "Tact Coa"
	for i := 0; i < b.N; i++ {
		isPalindromePermutation(input)
	}
}
