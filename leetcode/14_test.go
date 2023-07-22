package leetcode

import (
	"strings"
	"testing"
)

func longestCommonPrefix(strs []string) string {
	result := strings.Builder{}

	if len(strs) < 1 {
		return result.String()
	}

	for i := 0; i < len(strs[0]); i++ {
		compareStr := strs[0][i]
		same := true
		for _, value := range strs {
			if i >= len(value) || value[i] != compareStr {
				same = false
				break
			}
		}
		if !same {
			break
		}
		result.WriteByte(compareStr)
	}
	return result.String()
}

// TestLongestCommonPrefix tests the longestCommonPrefix function with various inputs.
func TestLongestCommonPrefix(t *testing.T) {
	testCases := []struct {
		input    []string
		expected string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{}, ""},
		{[]string{"", "abc", "def"}, ""},
		{[]string{"abc", "abc", "abc"}, "abc"},
	}

	for _, tc := range testCases {
		result := longestCommonPrefix(tc.input)
		if result != tc.expected {
			t.Errorf("Input: %v, Expected: %s, Got: %s", tc.input, tc.expected, result)
		}
	}
}

// BenchmarkLongestCommonPrefix benchmarks the longestCommonPrefix function.
func BenchmarkLongestCommonPrefix(b *testing.B) {
	strs := []string{"flower", "flow", "flight"}
	for i := 0; i < b.N; i++ {
		longestCommonPrefix(strs)
	}
}
