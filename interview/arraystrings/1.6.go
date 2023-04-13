package main

import (
	"strconv"
)

func compressString(s string) string {
	// If the length of the string is less than or equal to 2, it cannot be compressed
	if len(s) <= 2 {
		return s
	}

	compressed := make([]byte, 0, len(s)) // create a byte slice to hold compressed string

	var count int // count of consecutive characters
	for i := 0; i < len(s); i++ {
		count++ // increment count for each character

		// If the current character is different from the next character or if we have reached the end of the string,
		// add the current character and count to the compressed string
		if i+1 >= len(s) || s[i] != s[i+1] {
			compressed = append(compressed, s[i])
			compressed = append(compressed, []byte(strconv.Itoa(count))...)
			count = 0 // reset the count
		}
	}

	// If the compressed string is not smaller than the original string, return the original string
	if len(compressed) >= len(s) {
		return s
	}

	// Otherwise, return the compressed string
	return string(compressed)
}
