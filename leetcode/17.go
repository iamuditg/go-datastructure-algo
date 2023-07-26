package leetcode

func letterCombinations(digits string) []string {

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

	if len(digits) == 0 {
		return []string{}
	}
	var result []string
	solveCombination(digits, 0, "", &result, dict)
	return result
}

func solveCombination(digits string, index int, combination string, result *[]string, dict map[byte]string) {
	if index == len(digits) {
		*result = append(*result, combination)
		return
	}

	digit := digits[index]
	letter := dict[digit]

	for i := 0; i < len(letter); i++ {
		solveCombination(digits, index+1, combination+string(letter[i]), result, dict)
	}
}
