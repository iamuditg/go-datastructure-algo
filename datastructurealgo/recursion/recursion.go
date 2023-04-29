package main

import (
	"fmt"
)

func main() {
	//fmt.Println(factorial(5))
	//fmt.Println(fibonacci(3))
	//fmt.Println(power(5, 3))
	// printNum(6)
	//fmt.Println(numberOfDigit(12345))
	//fmt.Println(sumOfDigit(12345))
	//fmt.Println(multiplication(4, 3))
	//fmt.Println(countZeros(102030))
	//fmt.Println(geometricSum(3))
	//a := []int{1, 2, 3, 4, 5, 4, 8}
	//fmt.Println(checkArraySorted(a, len(a)))
	//fmt.Println(sumOfArray(a, len(a)))
	//fmt.Println(checkIfPresent(a, len(a), 8))
	//fmt.Println(firstIndexOfElement(a, len(a), 0, 4))
	//fmt.Println(lastIndexOfElement(a, len(a)-1, 0, 4))
	//printAllPosition(a, len(a)-1, 4)
	//fmt.Println(storeAllPosition(a, len(a)-1, 4, []int{}))
	str := "abvmba"
	fmt.Println(checkPalindrome(str, 0, len(str)-1))
}

func geometricSum(k int) float64 {
	if k == 0 {
		return 1
	}
	smallAns := geometricSum(k - 1)
	return smallAns + float64(1.0/power(2, k))
}

func countZeros(n int) int {
	if n == 0 {
		return 0
	}

	smallAns := countZeros(n / 10)

	last_digit := n % 10
	if last_digit == 0 {
		return smallAns + 1
	} else {
		return smallAns
	}
}

func multiplication(m int, n int) int {
	if n == 0 {
		return 0
	}
	smallAns := multiplication(m, n-1) + m
	return smallAns
}

func sumOfDigit(i int) int {
	if i == 0 {
		return 0
	}
	smallAns := sumOfDigit(i / 10)
	lastdigit := i % 10
	return smallAns + lastdigit
}

func numberOfDigit(i int) int {
	if i == 0 {
		return 0
	}
	smallAns := numberOfDigit(i / 10)
	return 1 + smallAns
}

func printNum(i int) {
	if i == 0 {
		return
	}
	printNum(i - 1)
	fmt.Print(i, "->")
}

func factorial(n int) int {
	if n == -1 {
		return -1
	}
	if n == 0 {
		return 1
	}
	smallFact := factorial(n - 1)
	return n * smallFact
}

func fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	smallOne := fibonacci(n - 1)
	smallTwo := fibonacci(n - 2)
	return smallOne + smallTwo
}

func power(x int, n int) int {
	if n == 0 {
		return 1
	}

	smallOutput := power(x, n-1)
	return x * smallOutput

}
