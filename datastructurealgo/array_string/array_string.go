package main

import "fmt"

func main() {

	//arr := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	//maxSum := MaxSubArraySum(arr)
	//fmt.Println("Maximum sum of contiguous subarray:", maxSum)

	//str := "A man, a plan, a canal: Panama"
	//isPal := IsPalindrome(str)
	//fmt.Println("Is palindrome?", isPal)

	//arr := []int{1, 7, 3, 6, 4, 6, 2, 8, 9}
	//rearranged := RearrangeArray(arr)
	//fmt.Println("Rearranged array:", rearranged)

	//str := "Hello World, how are you?"
	//reversed := ReverseWords(str)
	//fmt.Println("Reversed words:", reversed)

	//str1 := "Listenk"
	//str2 := "Silentk"
	//isAnagram := IsAnagram(str1, str2)
	//fmt.Println("Are the strings anagrams?", isAnagram)

	//str := "abcabcbabcd"
	//length := LengthOfLongestSubstring(str)
	//fmt.Println("Length of the longest substring:", length)

	str := "Hello World, how are you?"
	reversed := ReverseEachWords(str)
	fmt.Println("Reversed words:", reversed)

}
