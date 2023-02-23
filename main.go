package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	maxLen := 0
	charMap := make(map[byte]int)
	left := 0

	for right := 0; right < len(s); right++ {
		char := s[right]
		// if the current character is already in the window, move the left pointer
		// to the right of the previous occurrence
		if index, found := charMap[char]; found && index >= left {
			left = index + 1
		}
		// update the index of the current character in the map
		charMap[char] = right
		fmt.Println("left: ", left, " charMap: ", charMap)
		// update the maximum length seen so far
		if length := right - left + 1; length > maxLen {
			maxLen = length
			fmt.Println("maxLength: ", maxLen)
		}
	}
	return maxLen
}

func main() {
	fmt.Println(lengthOfLongestSubstring("ababcabcd"))
}
