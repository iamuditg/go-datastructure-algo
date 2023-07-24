// CODE EXAMPLE VALID FOR COMPILING
package main

import "fmt"

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	dict := map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}

	var result []string
	backtrack(digits, 0, "", dict, &result)
	return result
}

func backtrack(digits string, index int, combination string, dict map[byte]string, result *[]string) {
	if index == len(digits) {
		*result = append(*result, combination)
		return
	}

	letters := dict[digits[index]]
	for i := 0; i < len(letters); i++ {
		backtrack(digits, index+1, combination+string(letters[i]), dict, result)
	}
}

func main() {
	fmt.Println(letterCombinations("23"))
}
