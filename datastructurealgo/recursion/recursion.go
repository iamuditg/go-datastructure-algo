package main

import "fmt"

// -------------------- PROBLEM 1 -------------------------- //
// geometricSum calculates the sum of the geometric series
// 1 + 1/2 + 1/4 + ..... + 1/2^k
func geometricSum(k int) float64 {
	if k == 0 {
		return 1
	}
	smallAns := geometricSum(k - 1)
	return smallAns + float64(1.0/power(2, k))
}

// power calculates x raised to the power of n
func power(x int, n int) int {
	if n == 0 {
		return 1
	}
	smallOutput := power(x, n-1)
	return x * smallOutput
}

// -------------------- PROBLEM 2 -------------------------- //
// countZeros counts the number of zeros in the decimal representation
// of a non-negative integer
func countZeros(n int) int {
	fmt.Println("n - ", n)
	if n == 0 {
		return 0
	}
	smallAns := countZeros(n / 10)
	lastDigit := n % 10
	fmt.Println("Return value - ", smallAns, n, lastDigit)
	if lastDigit == 0 {
		return smallAns + 1
	} else {
		return smallAns
	}
}

// -------------------- PROBLEM 3 -------------------------- //
// multiplication performs multiplication of two integer without using the
// * operator.
func multiplication(m int, n int) int {
	fmt.Println("m : ", m, "n: ", n)
	if n == 0 {
		return 0
	}
	smallAns := multiplication(m, n-1) + m
	fmt.Println("smallAns: ", smallAns)
	return smallAns
}

// -------------------- PROBLEM 4 -------------------------- //
// sumOfDigit calculates the sum of digits in a given non-negative number
func sumOfDigit(i int) int {
	fmt.Println("i:- ", i)
	if i == 0 {
		return 0
	}
	lastDigit := i % 10
	i = i / 10
	output := sumOfDigit(i)
	fmt.Println("output: ", output, "lastDigit:", lastDigit)
	return output + lastDigit
}

// -------------------- PROBLEM 5 -------------------------- //
// numberOfDigit counts the number of digits in a given non-negative number
func numberOfDigit(i int) int {
	if i == 0 {
		return 0
	}
	smallAns := numberOfDigit(i / 10)
	return 1 + smallAns
}

// -------------------- PROBLEM 6 -------------------------- //
// printNum prints number from 1 to i in increasing order using recursion.
func printNum(i int) {
	if i == 0 {
		return
	}
	printNum(i - 1)
	fmt.Print(i, "->")
}

// -------------------- PROBLEM 7 -------------------------- //
// factorial calculates the factorial of a non-negative integer n.
func factorial(n int) int {
	fmt.Println("n:- ", n)
	if n == -1 {
		return -1
	}
	if n == 0 {
		return 1
	}
	ans := factorial(n - 1)
	fmt.Println("ans:- ", ans, " n:- ", n)
	return n * ans
}

// -------------------- PROBLEM 8 -------------------------- //
// fibonacci calculates the nth fibonacci number
func fibonacci(n int) int {
	fmt.Println("n: - ", n)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	// 0 , 1 , 1, 2, 3, 5
	smallOne := fibonacci(n - 1)
	fmt.Println("smallOne :- ", smallOne, " n: ", n)
	smallTwo := fibonacci(n - 2)
	fmt.Println("smallTwo :-", smallTwo, " n: ", n)
	fmt.Println("Output -", smallOne+smallTwo)
	return smallOne + smallTwo
}

// -------------------- PROBLEM 9 -------------------------- //
// checkArraySorted checks if the given array a is sorted in non-decreasing
// order.
func checkArraySorted(a []int, n int) bool {
	if n == 0 || n == 1 {
		return true
	}
	if a[0] > a[1] {
		return false
	}
	isSorted := checkArraySorted(a[1:], n-1)
	return isSorted
}

// -------------------- PROBLEM 10 -------------------------- //
// sumOfArray calculates the sum of the elements in the given array a.
func sumOfArray(a []int, n int) int {
	if n == 0 {
		return 0
	}
	sum := n + sumOfArray(a[1:], n-1)
	return sum
}

// -------------------- PROBLEM 11 -------------------------- //
// checkIfPresent check if the element is present in the given array a.
func checkIfPresent(a []int, n int, element int) bool {
	if n == 0 {
		return false
	}
	if a[0] == element {
		return true
	}
	check := checkIfPresent(a[1:], n-1, element)
	return check
}

// -------------------- PROBLEM 12 -------------------------- //
// firstIndexOfElement finds the first index of the element in the given array
func firstIndexOfElement(a []int, n int, i int, element int) int {
	if i == n {
		return -1
	}
	if a[i] == element {
		return i
	}
	index := firstIndexOfElement(a, n, i+1, element)
	return index
}

// -------------------- PROBLEM 13 -------------------------- //
// lastIndexOfElement finds the last index of the element in the given array
func lastIndexOfElement(a []int, n int, i int, element int) int {
	if n == i {
		return -1
	}
	if a[n] == element {
		return n
	}
	index := lastIndexOfElement(a, n-1, i, element)
	return index
}

// -------------------- PROBLEM 14 -------------------------- //
// printAllPosition prints all positions of the elements in the given array.
func printAllPosition(a []int, n int, element int) {
	if n < 0 {
		return
	}
	if a[n] == element {
		println(n)
	}
	printAllPosition(a, n-1, element)
}

// -------------------- PROBLEM 15 -------------------------- //
// storeAllPosition stores all position of the element in the given array
func storeAllPosition(a []int, n int, element int, store []int) []int {
	if n < 0 {
		return store
	}
	if a[n] == element {
		store = append(store, n)
	}
	store = storeAllPosition(a, n-1, element, store)
	return store
}

// -------------------- PROBLEM 16 -------------------------- //
// checkPalindrome checks if the given string str is a palindrome
func checkPalindrome(str string, i int, n int) bool {
	if i > n {
		return true
	}
	if str[i] != str[n] {
		return false
	}
	isPalindrome := checkPalindrome(str, i+1, n-1)
	return isPalindrome
}

// -------------------- PROBLEM 17 -------------------------- //
// printReverse prints the characters of the string in reverse order.
func printReverse(str string, n int) {
	fmt.Println(str, n)
	if n < 0 {
		return
	}
	printReverse(str[1:], n-1)
	println("After: ", string(str[0]), str)
	println(string(str[0]))
}

// -------------------- PROBLEM 18 -------------------------- //
// replaceCharacter replaces all occurrences of replaceChar in string n with x
func replaceCharacter(n string, replaceChar string) string {
	fmt.Println("n: ", n, replaceChar)
	if len(n) == 0 {
		return n
	}
	if string(n[0]) == replaceChar {
		return "x" + replaceCharacter(n[1:], replaceChar)
	} else {
		return string(n[0]) + replaceCharacter(n[1:], replaceChar)
	}
}

// -------------------- PROBLEM 19 -------------------------- //
// removeCharacter removes all occurrences of removeChar from string s.
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

// -------------------- PROBLEM 20 -------------------------- //
// removeConsecutiveDuplicates removes consecutive duplicate characters from the string.
func removeConsecutiveDuplicates(s string) string {
	if len(s) < 2 {
		return s
	}
	if s[0] == s[1] {
		return removeConsecutiveDuplicates(s[1:])
	}
	return string(s[0]) + removeConsecutiveDuplicates(s[1:])
}

// -------------------- PROBLEM 21 -------------------------- //
// printSubsequences prints all subsequences of the string.
func printSubsequences(str string) {
	output := ""
	printSubsequencesHelper(str, output, 0)
}

func printSubsequencesHelper(str string, output string, index int) {
	fmt.Println(str, output, index)
	if index == len(str) {
		if len(output) > 0 {
			fmt.Println(output)
		}
		return
	}

	// Exclude the current character
	printSubsequencesHelper(str, output, index+1)
	fmt.Println("Exclude", str, output, index)
	// Include the current character
	printSubsequencesHelper(str, output+string(str[index]), index+1)
	fmt.Println("Include", str, output, index)
}

func main() {
	//fmt.Println("PROBLEM 1 :- ",geometricSum(10))
	//fmt.Println("PROBLEM 2 :-", countZeros(10102030))
	//fmt.Println("PROBLEM 3 :- ",multiplication(3, 4))
	//fmt.Println("PROBLEM 4 :- ", sumOfDigit(12345))
	//fmt.Println("PROBLEM 5 :- ", numberOfDigit(12345))
	//fmt.Println("PROBLEM 6 :- ")
	//printNum(6)
	//fmt.Println("PROBLEM 7 :- ", factorial(5))
	//fmt.Println("PROBLEM 8 :- ", fibonacci(5))
	//fmt.Println("PROBLEM 9 :- ", checkArraySorted([]int{1, 3, 8, 4}, 4))
	//fmt.Println("PROBLEM 10 :- ", sumOfArray([]int{1, 2, 3}, 3))
	//fmt.Println("PROBLEM 11 :- ", checkIfPresent([]int{3, 9, 5, 2, 4, 1}, 6, 2))
	//fmt.Println("PROBLEM 12 :- ", firstIndexOfElement([]int{1, 2, 2, 3, 4, 4, 5, 6}, 7, 0, 4))
	//fmt.Println("PROBLEM 13 :- ", lastIndexOfElement([]int{1, 2, 2, 3, 4, 4, 5, 6}, 7, 0, 4))
	//fmt.Println("PROBLEM 14 :-")
	//printAllPosition([]int{2, 3, 4, 5, 4, 6, 4}, 6, 4)
	//fmt.Println("PROBLEM 15 :- ", storeAllPosition([]int{1, 2, 2, 3, 4, 4, 5, 6}, 7, 4, []int{}))
	//fmt.Println("PROBLEM 16 :- ", checkPalindrome("aabbaa", 0, 5))
	//fmt.Println("PROBLEM 17 :- ")
	//printReverse("abcd", 3)
	//fmt.Println("PROBLEM 18 :- ", replaceCharacter("abccdcd", "c"))
	//fmt.Println("PROBLEM 19 :- ", removeCharacter("abccdcd", "c"))
	//fmt.Println("PROBLEM 20 :- ", removeConsecutiveDuplicates("aabbccdd"))
	fmt.Println("PROBLEM 21 :- ")
	printSubsequences("abc")
}
