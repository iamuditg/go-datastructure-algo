package main

import "fmt"

func letterCombinations(digits string) []string {
	var result []string
	if digits == "" {
		return result
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

	result = []string{""}

	for i := range digits {
		st := dict[digits[i]]
		fmt.Println("st:", string(st))
		var next []string
		for j := range st {
			c := st[j]
			fmt.Println("c: ", string(c))
			for _, r := range result {
				fmt.Println("r:", string(r))
				next = append(next, r+string(c))
				fmt.Println("next: ", next)
			}
		}
		result = next
		fmt.Println("result:", result)
	}
	return result
}

func main() {
	fmt.Println(letterCombinations("23"))
}
