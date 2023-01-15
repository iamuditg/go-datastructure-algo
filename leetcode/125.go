package leetcode

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return 'a' + c - 'A'
	}
	return c
}

func isAlphanumeric(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func isPalindrome(s string) bool {
	i, j := 0, len(s)-1
	for i < j {
		left, right := s[i], s[j]
		for !isAlphanumeric(left) && i < j {
			i++
			left = s[i]
		}
		for !isAlphanumeric(right) && i < j {
			j--
			right = s[j]
		}
		if toLower(left) != toLower(right) {
			return false
		}
		i++
		j--
	}
	return true
}
