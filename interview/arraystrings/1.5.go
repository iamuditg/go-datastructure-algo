package main

func oneEditAway(s1, s2 string) bool {
	// if the length of the string differ by more than 1, they cannot be one edit
	if abs(len(s1)-len(s2)) > 1 {
		return false
	}

	// iterate over both strings at the same time, checking if each character matches
	i, j, edits := 0, 0, 0
	for i < len(s1) && j < len(s2) {
		if s1[i] != s2[j] {
			// If the character don't match, increment the number of edits
			edits++
			if edits > 1 {
				// If the number of edits exceeds 1, the strings are not one edit away
				return false
			}
			if len(s1) == len(s2) {
				// If the strings have the same length , we can replace the character
				i++
				j++
			} else if len(s1) < len(s2) {
				// If s2 is longer than s1, we can insert a character into s1
				j++
			} else {
				// If s1 is longer than s2, we can remove a character from s1
				i++
			}
		} else {
			// if the character match , move on the next character in both string
			i++
			j++
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
