package leetcode

import "math"

func divide(dividend int, divisor int) int {
	// check for overflow and special case
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		return -dividend
	}

	// determine sign of the quotient
	sign := 1
	if dividend < 0 {
		sign = -sign
		dividend = -dividend
	}
	if divisor < 0 {
		sign = -sign
		divisor = -divisor
	}

	// Initialize variables
	quotient := 0
	power := 31

	// find the highest power of divisor that is less thean or equal to divident
	for dividend >= divisor {
		if divisor<<power <= dividend {
			quotient += 1 << power
			dividend -= divisor << power
		}
		power -= 1
	}

	// apply sign to quotient and return
	if sign == -1 {
		return -quotient
	}

	return quotient
}
