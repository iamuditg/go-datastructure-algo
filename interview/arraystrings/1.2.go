package main

/*
Check Permutation: Given two strings, write a method to decide if one is a permutation of the
other.
*/

func isPermutation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	// count the number of occurrence of each character is s1
	charCount := make(map[rune]int)
	for _, char := range s1 {
		charCount[char]++
	}

	// Decrement the count of each character is s2v
	for _, char := range s2 {
		charCount[char]--
		if charCount[char] < 0 {
			return false
		}
	}

	return true

}
