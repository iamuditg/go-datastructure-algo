package main

import "testing"

/*
Check Permutation: Given two strings, write a method to decide if one is a permutation of the
other.
*/

func isPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	// Create a character count map for s1
	charCount := make(map[rune]int)

	// Count the occurrences of each character in s1
	for _, char := range s1 {
		charCount[char]++
	}

	// Decrement the count of each character in s2
	// If any character count becomes negative or a character is missing in s1, return false
	for _, char := range s2 {
		charCount[char]--
		if charCount[char] < 0 {
			return false
		}
	}
	return true
}

func TestIsPermutation(t *testing.T) {
	tests := []struct {
		s1       string
		s2       string
		expected bool
	}{
		{"abc", "cba", true},
		{"abc", "def", false},
		{"aab", "abb", false},
		{"aabbcc", "", false},
		{"", "", true},
	}

	for _, test := range tests {
		result := isPermutation(test.s1, test.s2)
		if result != test.expected {
			t.Errorf("isPermutation(%s, %s) = %t, expected %t", test.s1, test.s2, result, test.expected)
		}
	}
}

func BenchmarkIsPermutation(b *testing.B) {
	s1 := "abcdefghijklmnopqrstuvwxyz"
	s2 := "zyxwvutsrqponmlkjihgfedcba"
	for i := 0; i < b.N; i++ {
		isPermutation(s1, s2)
	}
}
