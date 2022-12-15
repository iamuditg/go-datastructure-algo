package leetcode

import (
	"fmt"
	"strings"
)

func longestCommonPrefix(strs []string) string {
	result := strings.Builder{}
	if len(strs) == 0 {
		return result.String()
	}
	fmt.Println(strs[0])
	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		fmt.Println("s:", string(c))
		same := true
		for _, s := range strs {
			fmt.Println(string(s[i]))
			if i >= len(s) || s[i] != c {
				same = false
				break
			}
		}
		if !same {
			break
		}
		result.WriteByte(c)
	}
	return result.String()
}
