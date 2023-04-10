package main

import "fmt"

func hasUniqueChars(str string) bool {
	charSet := make(map[rune]bool)
	for i, char := range str {
		fmt.Println(i, char)
		if charSet[char] {
			return false
		}
		charSet[char] = true
	}
	return true
}

func hasUniqueCharsOtherDS(str string) bool {
	for i, char := range str {
		for j := i + 1; j < len(str); j++ {
			fmt.Println(char, str[j])
			if char == rune(str[j]) {
				return false
			}
		}
	}
	return true
}
