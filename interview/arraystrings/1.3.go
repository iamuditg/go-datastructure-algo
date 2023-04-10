package main

import "fmt"

func URLify(s string, length int) string {
	spaceCount := 0
	// Count the number of spaces in the string
	for i := 0; i < length; i++ {
		if s[i] == ' ' {
			spaceCount++
		}
	}
	fmt.Println(spaceCount)
	// calculate the new length of the string after replacing spaces with "%20"
	newLength := length + spaceCount*2
	fmt.Println(newLength)
	// create a new byte slice of the appropriate length to hold the modifies string
	newS := make([]byte, newLength)

	j := 0
	//  copy each character to its new position in the string, replacing spaces wih "%20"
	for i := 0; i < length; i++ {
		if s[i] == ' ' {
			newS[j] = '%'
			newS[j+1] = '2'
			newS[j+2] = '0'
			j += 3
		} else {
			newS[j] = s[i]
			j++
		}
	}
	fmt.Println(string(newS))

	// Remove any trailing spaces from the new string
	// iterate over the slices from the end and decrement the new length until a non-space character is found
	for i := newLength - 1; i >= 0; i-- {
		if newS[i] == ' ' {
			newLength--
		} else {
			break
		}
	}

	// convert the byte slice to a string and return it
	return string(newS[:newLength])
}
