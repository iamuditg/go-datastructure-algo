package leetcode

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
		searchDict := dict[digits[i]]
		var next []string
		for j := range searchDict {
			word := searchDict[j]
			for _, r := range result {
				next = append(next, r+string(word))
			}
		}
		result = next
	}
	return result
}
