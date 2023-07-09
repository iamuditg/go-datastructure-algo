package main

import (
	"strconv"
	"testing"
)

/*
String Compression:
Implement a method to perform basic string compression using the counts
of repeated characters. For example, the string aabcccccaaa would become a2blc5a3. If the
"compressed" string would not become smaller than the original string, your method should return
the original string. You can assume the string has only uppercase and lowercase letters (a - z).
*/

func compressString(s string) string {
	if len(s) <= 2 {
		return s
	}

	compressed := make([]byte, 0, len(s))
	var count int
	for i := 0; i < len(s); i++ {
		count++

		if i+1 >= len(s) || s[i] != s[i+1] {
			compressed = append(compressed, s[i])
			compressed = append(compressed, []byte(strconv.Itoa(count))...)
			count = 0
		}
	}
	if len(compressed) >= len(s) {
		return s
	}

	return string(compressed)
}

func TestCompressString(t *testing.T) {
	tests := []struct {
		s        string
		expected string
	}{
		{"aabcccccaaa", "a2b1c5a3"},
		{"abc", "abc"},
		{"aabbcc", "aabbcc"},
		{"aaaaa", "a5"},
		{"", ""},
	}

	for _, test := range tests {
		result := compressString(test.s)
		if result != test.expected {
			t.Errorf("compressString(%q) = %q, expected %q", test.s, result, test.expected)
		}
	}
}

func BenchmarkCompressString(b *testing.B) {
	s := "aabcccccaaa"
	for i := 0; i < b.N; i++ {
		compressString(s)
	}
}
