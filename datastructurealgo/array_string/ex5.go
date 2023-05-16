package main

import (
	"fmt"
	"strings"
)

func IsAnagram(str1, str2 string) bool {
	comp := make(map[rune]int)
	// to lower case
	str1 = strings.ToLower(str1)
	str2 = strings.ToLower(str2)

	for _, val := range str1 {
		if _, ok := comp[val]; !ok {
			comp[val]++
		} else {
			comp[val]--
		}
	}

	for _, val := range str2 {
		if _, ok := comp[val]; !ok {
			comp[val]++
		} else {
			comp[val]--
		}
	}

	fmt.Println(comp)

	for _, v := range comp {
		if v >= 1 {
			return false
		}
	}

	return true
}
