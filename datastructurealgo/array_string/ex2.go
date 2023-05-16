package main

import "strings"

/*
	Given a string, determine if it is palindrome( reads the same forwards and backwards ) after removing non-alphanumeric characters and ignoring case.
*/

func IsPalindrome(s string) bool {
	normalized := normalizeString(s)

	i := 0
	j := len(normalized) - 1

	for i < j {
		if normalized[i] != normalized[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func normalizeString(s string) string {
	// Remove non-alphanumeric characters and convert to lowercase
	var sb strings.Builder

	for _, ch := range s {
		if isAlphanumeric(ch) {
			sb.WriteRune(toLower(ch))
		}
	}
	return sb.String()
}

func toLower(ch rune) rune {
	if ch >= 'A' && ch <= 'Z' {
		return ch + ('a' - 'A')
	}
	return ch
}

func isAlphanumeric(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9')
}
