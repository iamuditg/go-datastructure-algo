package leetcode

import (
	"math"
	"testing"
)

func myAtoi(str string) int {

	s := 0
	// check str contains only numeric value
	for s = 0; s < len(str); s++ {
		if !isDigitOrSignValue(str[s]) {
			return 0
		}

		if isSignValue(str[s]) {
			break
		}
	}

	var position bool
	if str[s] == '-' {
		position = false
		s++
	} else if str[s] == '+' {
		position = true
		s++
	} else {
		position = true
	}

	result := 0

	for ; s < len(str); s++ {
		if !isDigit(str[s]) {
			return result
		}
		value := int(str[s]) - '0'
		if !position {
			value = -value
		}

		nextValue := result*10 + value

		if nextValue > math.MaxInt32 || nextValue < math.MinInt32 || (nextValue-value)/10 != result {
			if position {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}

		result = nextValue

	}

	return result
}

func isDigitOrSignValue(digit byte) bool {
	if isSignValue(digit) || digit == ' ' {
		return true
	} else {
		return false
	}
}

func isSignValue(digit byte) bool {
	if isDigit(digit) || digit == '+' || digit == '-' {
		return true
	} else {
		return false
	}
}

func isDigit(digit byte) bool {
	if digit >= '0' && digit <= '9' {
		return true
	} else {
		return false
	}
}

// TestMyAtoi tests the myAtoi function with various inputs.
func TestMyAtoi(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{"42", 42},
		{"   -42", -42},
	}

	for _, tc := range testCases {
		result := myAtoi(tc.input)
		if result != tc.expected {
			t.Errorf("Input: %s, Expected: %d, Got: %d", tc.input, tc.expected, result)
		}
	}
}

// BenchmarkMyAtoi benchmarks the myAtoi function.
func BenchmarkMyAtoi(b *testing.B) {
	for i := 0; i < b.N; i++ {
		myAtoi("42")
	}
}
