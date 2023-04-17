package main

import "fmt"

func isRotation(s1 string, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	// find the first character of ss1 in s2
	found := false
	for i := 0; i < len(s2); i++ {
		fmt.Println(string(s1[0]), string(s2[i]))
		if s1[0] == s2[i] {
			// check if the rest of s1 matches the character in s2
			// starting from the current position
			matches := true
			for j := 0; j < len(s1); j++ { //s1 := "waterbottle" s2 := "erbottlewat"
				fmt.Println("inside: ", string(s1[j]), string(s2[(i+j)%len(s2)]), i+j, len(s2), (i+j)%len(s2))
				if s1[j] != s2[(i+j)%len(s2)] {
					matches = false
					break
				}
			}
			if matches {
				found = true
				break
			}
		}
	}
	return found
}
