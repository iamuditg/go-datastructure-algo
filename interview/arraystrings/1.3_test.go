package main

import "testing"

/*
URLify:
Write a method to replace all spaces in a string with '%20'. You may assume that the string
has sufficient space at the end to hold the additional characters, and that you are given the "true"
length of the string.
EXAMPLE
Input: "Mr John Smith ", 13
Output: "Mr%20John%20Smith"
*/

func URLify(s string, length int) string {
	spaceCount := 0

	// count the number of spaces in the string
	for i := 0; i < length; i++ {
		if s[i] == ' ' {
			spaceCount++
		}
	}

	// Calculate the new length of the string after replacing spaces with "%20"
	newLength := length + spaceCount*2

	// Create a new byte slice of the appropriate length to hold the modified string
	newS := make([]byte, newLength)

	j := 0
	// Copy each character to its new position in the string, replacing space with "%20"
	for i := 0; i < length; i++ {
		if s[i] == ' ' {
			newS[j] = '%'
			newS[j+1] = '2'
			newS[j+2] = '0'
			j += 3
		} else {
			newS[j] = s[i]
			j++
		}
	}

	// Remove any trailing spaces from the new string
	// Iterate over the slice from the end and decrement the new length until a non-space character is found
	for i := newLength - 1; i >= 0; i-- {
		if newS[i] == ' ' {
			newLength--
		} else {
			break
		}
	}

	return string(newS[:newLength])
}

func TestURLify(t *testing.T) {
	tests := []struct {
		input    string
		length   int
		expected string
	}{
		{"Mr John Smith ", 13, "Mr%20John%20Smith"},
		{"Hello World", 11, "Hello%20World"},
		{"     ", 2, "%20%20"},
	}

	for _, test := range tests {
		result := URLify(test.input, test.length)
		if result != test.expected {
			t.Errorf("URLify(%q, %d) = %q, expected %q", test.input, test.length, result, test.expected)
		}
	}
}

func BenchmarkURLify(b *testing.B) {
	input := "Mr John Smith " // Length = 13
	length := 13
	for i := 0; i < b.N; i++ {
		URLify(input, length)
	}
}
