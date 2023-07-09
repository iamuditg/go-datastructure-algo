package main

import "testing"

/*
One Away: There are three types of edits that can be performed on strings: insert a character,
remove a character, or replace a character. Given two strings, write a function to check if they are
one edit (or zero edits) away.
EXAMPLE
pale, ple -> true
pales, pale -> true
pale, bale -> true
pale, bake -> false
*/

func oneEditAway(s1, s2 string) bool {
	if abs(len(s1)-len(s2)) > 1 {
		return false
	}

	edits := 0
	i, j := 0, 0

	for i < len(s1) && j < len(s2) {
		if s1[i] != s2[j] {
			if edits == 1 {
				return false
			}

			edits++
			if len(s1) < len(s2) {
				j++
			} else if len(s1) > len(s2) {
				i++
			} else {
				i++
				j++
			}
		} else {
			i++
			j++
		}
	}

	if i < len(s1) || j < len(s2) {
		edits++
	}

	return edits <= 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func TestOneEditAway(t *testing.T) {
	tests := []struct {
		s1       string
		s2       string
		expected bool
	}{
		{"pale", "ple", true},
		{"pales", "pale", true},
		{"pale", "bale", true},
		{"pale", "bake", false},
		{"hello", "hello", true},
		{"abcd", "abcd", true},
		{"abcdef", "abefcd", false},
		{"abcdef", "abcfed", true},
	}

	for _, test := range tests {
		result := oneEditAway(test.s1, test.s2)
		if result != test.expected {
			t.Errorf("oneEditAway(%q, %q) = %v, expected %v", test.s1, test.s2, result, test.expected)
		}
	}
}

func BenchmarkOneEditAway(b *testing.B) {
	s1 := "pale"
	s2 := "bake"
	for i := 0; i < b.N; i++ {
		oneEditAway(s1, s2)
	}
}
