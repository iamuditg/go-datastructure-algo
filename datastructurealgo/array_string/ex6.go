package main

import "fmt"

/*
	Given a string, find the length of the longest substring without repeating characters.
*/

func LengthOfLongestSubstring(s string) int {
	seen := make(map[byte]int)
	maxLength := 0
	start := 0

	for i := 0; i < len(s); i++ {
		if idx, ok := seen[s[i]]; ok && idx >= start {
			start = idx + 1
		}
		seen[s[i]] = i
		currLength := i - start + 1
		fmt.Println(seen, maxLength, start, currLength)
		if currLength > maxLength {
			maxLength = currLength
		}
	}
	return maxLength
}
