package main

import "strings"

func ReverseEachWords(s string) string {
	words := strings.Fields(s)

	for i := 0; i < len(words); i++ {
		words[i] = reverseString(words[i])
	}
	return strings.Join(words, " ")
}

func reverseString(s string) string {
	runes := []rune(s)
	i := 0
	j := len(runes) - 1

	for i < j {
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}

	return string(runes)
}
