package main

import "fmt"

func printReverse(str string, n int) {
	if n < 0 {
		return
	}
	printReverse(str[1:], n-1)
	fmt.Println(string(str[0]))
}

// -------------------------------------------------------
func findLength(n []int) int {
	if n == nil && n[1:] != nil {
		return 1
	}
	smallAnswer := findLength(n[1:])
	smallAnswer += 1
	return smallAnswer
}

// --------------------------------------------------------
func replaceCharacter(n string, replaceChar string) string {
	if len(n) == 0 {
		return n
	}
	fmt.Println(n, n[0])
	if string(n[0]) == replaceChar {
		return "x" + replaceCharacter(n[1:], replaceChar)
	} else {
		return string(n[0]) + replaceCharacter(n[1:], replaceChar)
	}
}

// ------------------------------------------------------------
func removeCharacter(s string, removeChar string) string {
	if len(s) == 0 {
		return s
	}

	if string(s[0]) == removeChar {
		return removeCharacter(s[1:], removeChar)
	} else {
		return string(s[0]) + removeCharacter(s[1:], removeChar)
	}
}

// --------------------------------------------------------
func removeConsecutiveDuplicates(s string) string {
	if len(s) < 2 {
		return s
	}
	if s[0] == s[1] {
		return removeConsecutiveDuplicates(s[1:])
	}
	return string(s[0]) + removeConsecutiveDuplicates(s[1:])
}

// --------------------------------------------------------
func printSubsequences(str string) {
	output := ""
	printSubsequencesHelper(str, output, 0)
}

func printSubsequencesHelper(str string, output string, index int) {
	if index == len(str) {
		if len(output) > 0 {
			fmt.Println(output)
		}
		return
	}

	// Exclude the current character
	fmt.Println("Exclude:", str, output, index)
	printSubsequencesHelper(str, output, index+1)

	// Include the current character
	fmt.Println("Include:", str, output, index)
	printSubsequencesHelper(str, output+string(str[index]), index+1)
}

// ----------------------------------------------------------------------
func stringToInt(str string) int {
	return stringToIntRecursive(str, len(str))
}

func stringToIntRecursive(str string, index int) int {
	if index == 0 {
		return 0
	}
	smallAns := stringToIntRecursive(str, index-1)
	lastDigit := str[index-1] - '0'
	ans := smallAns*10 + int(lastDigit)
	fmt.Println(smallAns, index, str[index-1], '0', lastDigit, ans)
	return ans
}

//-----------------------------------------------------------------------
