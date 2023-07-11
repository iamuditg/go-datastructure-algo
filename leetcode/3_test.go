package leetcode

import "testing"

// Problem - https://leetcode.com/problems/longest-substring-without-repeating-characters/

// lengthOfLongestSubstring calculates the length of the longest substring without repeating characters.
func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	charMap := make(map[byte]int)
	left := 0

	for right := 0; right < len(s); right++ {
		char := s[right]

		// If the current character is already in the window, move the left pointer
		// to the right of the previous occurrence
		if index, found := charMap[char]; found && index >= left {
			left = index + 1
		}

		// Update the index of the current character in the map
		charMap[char] = right

		// Update the maximum length seen so far
		if length := right - left + 1; length > maxLen {
			maxLen = length
		}
	}
	return maxLen
}

func TestLengthOfLongestSubstring(t *testing.T) {
	tests := []struct {
		s        string
		expected int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"", 0},
	}

	for _, test := range tests {
		result := lengthOfLongestSubstring(test.s)
		if result != test.expected {
			t.Errorf("lengthOfLongestSubstring(%s) = %d, expected %d", test.s, result, test.expected)
		}
	}
}

func BenchmarkLengthOfLongestSubstring(b *testing.B) {
	s := "abcabcbb"
	for i := 0; i < b.N; i++ {
		lengthOfLongestSubstring(s)
	}
}
