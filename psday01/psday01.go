package psday01

import (
	"fmt"
	"strings"
)

/*
1)	Given a string, find the words with the maximum number of vowels.

Input: "The quick brown fox jumps right over the little lazy dog."
		Maximum Number of Vowels: 2
	output:	quick, over, little
*/

func FindMaximumNumberOfWords(str string, max int) []string {
	count, index := 0, 0
	output := make([]string, 0)
	for i := 0; i < len(str); i++ {
		switch str[i] {
		case ' ':
			if count == max {
				result := strings.Trim(string(str[index:i]), " ")
				output = append(output, result)
			}
			index = i
			count = 0
			break
		case 'a', 'e', 'i', 'o', 'u':
			count++
		}
	}
	return output
}

/*
2)	Find the next greatest element for each element in the Array
Input: {15, 10, 16, 20, 8, 9, 7, 50}

Output:
 15 -16
 10 -16
 16-20
 20-50
  8-9
  9-50
  7-50
  50 -1
*/

func FindGreatestElementArray(arr []int) map[int]int {
	output := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		output[arr[i]] = -1
		for j := i; j < len(arr); j++ {
			if arr[i] < arr[j] {
				output[arr[i]] = arr[j]
				break
			}
		}
	}
	return output
}

/*
6)	Program to check if the input string has valid parentheses. A parenthesis is valid if
o	Open parentheses should be closed in the same order. Meaning that the below will be invalid as parentheses are not closed in right order ([)]
o	Every open parenthesis should have the corresponding close parentheses
Input: {}
Output: valid

Input {(})
Output: invalid

*/

func CheckValidParenthesis(str string) bool {
	if len(str)%2 != 0 {
		return false
	}
	for i := 1; i < len(str); i += 2 {
		if str[i-1] == '{' && str[i] == '}' {
			continue
		} else {
			return false
		}
	}
	return true
}

/*
5) Write a program to check whether the given string is palindrome or not? If not, then please check whether by removing only one char from the string we can make it palindrome.
Input: abbad
Output: is a palindrome by removing last char “d”, value: abba

Input: badd
Output: is not a palindrome
*/

func CheckPalindromeString(str string) string {
	left := 0
	right := len(str) - 1
	output := ""
	for left < right {
		if str[left] != str[right] {
			leftIncr := check(str, left+1, right)
			rightIncr := check(str, left, right-1)
			if rightIncr {
				output = string(str[right])
			} else if leftIncr {
				output = string(str[left])
			}
			if output == "" {
				return fmt.Sprintf("%s is not palindrome", str)
			} else {
				return fmt.Sprintf("is a palindrome by removing last char \"%s\" , value: %s", output, str)
			}
		} else {
			left++
			right--
		}
	}

	return fmt.Sprintf("%s is palindrome", str)
}

func check(str string, left int, right int) bool {
	for left < right {
		if str[left] != str[right] {
			return false
		}
		left++
		right--
	}
	return true
}
