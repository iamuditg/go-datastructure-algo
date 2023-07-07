package main

import "testing"

/* PROBLEM  -
Is Unique: Implement an algorithm to determine if a string has all unique characters. What if you
cannot use additional data structures?
*/

// SOLUTION -

// hasUniqueChars checks if a string has all unique characters using a hash set.
func hasUniqueChars(str string) bool {
	charSet := make(map[rune]bool)
	for _, char := range str {
		if charSet[char] {
			return false
		}
		charSet[char] = true
	}
	return true
}

// hasUniqueCharsOtherDS checks if a string has all unique characters without using additional data structures.
// It uses nested loops to compare each character with all subsequent characters.
func hasUniqueCharsOtherDS(str string) bool {
	for i, char := range str {
		for j := i + 1; j < len(str); j++ {
			if char == rune(str[j]) {
				return false
			}
		}
	}
	return true
}

func TestHasUniqueChars(t *testing.T) {
	testCases := []struct {
		str      string
		expected bool
	}{
		{"abcdefg", true},
		{"hello", false},
		{"12345", true},
		{"", true},
	}

	for _, tc := range testCases {
		result := hasUniqueChars(tc.str)
		if result != tc.expected {
			t.Errorf("hasUniqueChars failed for %s. Expected: %t, Got: %t", tc.str, tc.expected, result)
		}
	}
}

func TestHasUniqueCharsOtherDS(t *testing.T) {
	testCases := []struct {
		str      string
		expected bool
	}{
		{"abcdefg", true},
		{"hello", false},
		{"12345", true},
		{"", true},
	}

	for _, tc := range testCases {
		result := hasUniqueCharsOtherDS(tc.str)
		if result != tc.expected {
			t.Errorf("hasUniqueCharsOtherDS failed for %s. Expected: %t, Got: %t", tc.str, tc.expected, result)
		}
	}
}

func BenchmarkHasUniqueChars(b *testing.B) {
	str := "abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < b.N; i++ {
		hasUniqueChars(str)
	}
}

func BenchmarkHasUniqueCharsOtherDS(b *testing.B) {
	str := "abcdefghijklmnopqrstuvwxyz"
	for i := 0; i < b.N; i++ {
		hasUniqueCharsOtherDS(str)
	}
}
