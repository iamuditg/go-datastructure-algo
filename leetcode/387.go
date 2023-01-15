package leetcode

import "fmt"

func firstUniqueChar(s string) int {
	dic := make([]int, 26)
	for i := 0; i < len(s); i++ {
		dic[s[i]-'a'] += 1
	}
	fmt.Println(dic)
	for i := 0; i < len(s); i++ {
		if dic[s[i]-'a'] == 1 {
			return i
		}
	}
	return -1
}
