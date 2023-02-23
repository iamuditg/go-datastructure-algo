package leetcode

func strStr(haystack string, needle string) int {
	// Handle edge cases
	if needle == "" {
		return 0
	}
	if len(haystack) < len(needle) {
		return -1
	}
	// loop through the haystack string
	for i := 0; i <= len(haystack)-len(needle); i++ {
		// compare each substring of the sum length as needle with needle
		match := true
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				match = false
				break
			}
		}
		// if we find a match, return the index of the starting character
		if match {
			return i
		}
	}
	return -1
}
