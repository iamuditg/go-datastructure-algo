package leetcode

import "fmt"

func isPowerOfThree(n int) bool {
	if n == 1 {
		return true
	}
	if n == 0 || n%3 != 0 {
		return false
	}
	for n != 0 {
		if n == 3 {
			return true
		}
		if n%3 == 0 {
			n /= 3
			fmt.Println(n)
		} else {
			return false
		}
	}
	return true
}
