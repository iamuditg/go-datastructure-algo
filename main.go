// CODE EXAMPLE VALID FOR COMPILING
package main

import (
	"fmt"
	"math"
)

func myAtoi(str string) int {

	s := 0
	// check str contains only numeric value
	for s = 0; s < len(str); s++ {
		if !isDigitOrSignValue(str[s]) {
			return 0
		}

		if isSignValue(str[s]) {
			break
		}
	}

	var position bool
	if str[s] == '-' {
		position = false
		s++
	} else if str[s] == '+' {
		position = true
		s++
	} else {
		position = true
	}

	result := 0

	for ; s < len(str); s++ {
		if !isDigit(str[s]) {
			return result
		}
		value := int(str[s]) - '0'
		if !position {
			value = -value
		}

		nextValue := result*10 + value

		if nextValue > math.MaxInt32 || nextValue < math.MinInt32 || (nextValue-value)/10 != result {
			if position {
				return math.MaxInt32
			} else {
				return math.MinInt32
			}
		}

		result = nextValue

	}

	return result
}

func isDigitOrSignValue(digit byte) bool {
	if isSignValue(digit) || digit == ' ' {
		return true
	} else {
		return false
	}
}

func isSignValue(digit byte) bool {
	if isDigit(digit) || digit == '+' || digit == '-' {
		return true
	} else {
		return false
	}
}

func isDigit(digit byte) bool {
	if digit >= '0' && digit <= '9' {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(myAtoi("   -42"))
}
