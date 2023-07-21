package leetcode

import (
	"math"
	"testing"
)

// Problem - https://leetcode.com/problems/reverse-integer/
func reverse(x int) int {
	ans := 0
	for x != 0 {
		last := x % 10
		x /= 10
		next := ans*10 + last

		// Check for integer overflow or underflow
		if (next-last)/10 != ans || next > math.MaxInt32 || next < math.MinInt32 {
			return 0
		}
		ans = next
	}
	return ans
}

// TestReverse tests the reverse function with various inputs.
func TestReverse(t *testing.T) {
	testCases := []struct {
		input    int
		expected int
	}{
		{123, 321},
		{-123, -321},
		{120, 21},
		{0, 0},
		{2147483647, 0},            // Exceeds int32 max value, expect 0
		{-2147483648, 0},           // Exceeds int32 min value, expect 0
		{7463847412, 2147483647},   // Exceeds int32 max value, expect 0
		{-9463847412, -2147483649}, // Exceeds int32 min value, expect 0
	}

	for _, tc := range testCases {
		result := reverse(tc.input)
		if result != tc.expected {
			t.Errorf("Input: %d, Expected: %d, Got: %d", tc.input, tc.expected, result)
		}
	}
}

// BenchmarkReverse benchmarks the reverse function.
func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverse(123456789)
	}
}
