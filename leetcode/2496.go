package leetcode

func maximumValue(strs []string) int {
	maxVal := 0
	for _, str := range strs {
		if isNumeric(str) {
			num := convertToInt(str)
			if num > maxVal {
				maxVal = num
			}
		} else {
			length := len(str)
			if length > maxVal {
				maxVal = length
			}
		}
	}
	return maxVal
}

func isNumeric(str string) bool {
	for _, ch := range str {
		if ch < '0' || ch > '9' {
			return false
		}
	}
	return true
}

func convertToInt(str string) int {
	num := 0
	for _, ch := range str {
		digit := int(ch - '0')
		num = num*10 + digit
	}
	return num
}
