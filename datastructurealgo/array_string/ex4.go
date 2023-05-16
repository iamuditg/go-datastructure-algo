package main

import (
	"fmt"
	"strings"
)

func ReverseWords(s string) string {
	words := strings.Fields(s)
	fmt.Println(words)
	reverseSlice(words)
	return strings.Join(words, " ")
}

func reverseSlice(s []string) {
	i := 0
	j := len(s) - 1
	for i < j {
		fmt.Println(s[i], s[j])
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}
