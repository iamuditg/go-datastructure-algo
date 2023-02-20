package leetcode

import "math"

func myAtoi(str string) int {
	s := 0
	for s = 0; s < len(str); s++ {
		if !isDigitOrSignOrWhiteSpace(str[s]) {
			return 0
		}

		// find first +/- or digit
		if isDigitOrSign(str[s]) {
			break
		}
	}

	var position bool
	if s < len(str) {
		if str[s] == '-' {
			position = false
			s++
		} else if str[s] == '+' {
			s++
			position = true
		} else {
			position = true
		}
	}

	result := 0
	for ; s < len(str); s++ {
		if !isDigit(str[s]) {
			break
		}

		d := int(str[s]) - '0'
		var nextResult int
		if !position {
			d = -d
		}

		nextResult = result*10 + d

		if nextResult > math.MaxInt32 || nextResult < math.MinInt32 || (nextResult-d)/10 != result {
			if position {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}
		result = nextResult
	}

	return result
}

func isDigitOrSignOrWhiteSpace(c byte) bool {
	if isDigitOrSign(c) || c == ' ' {
		return true
	} else {
		return false
	}
}

func isDigitOrSign(c byte) bool {
	if isDigit(c) || c == '-' || c == '+' {
		return true
	} else {
		return false
	}
}

func isDigit(c byte) bool {
	if c >= '0' && c <= '9' {
		return true
	} else {
		return false
	}
}
