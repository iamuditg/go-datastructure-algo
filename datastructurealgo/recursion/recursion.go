package main

import "fmt"

func countDown(number int) {
	fmt.Println(number)
	if number == 1 {
		return
	}
	countDown(number - 1)
}

// program to get factorial of number
func factorial(num int) int {
	if num == 1 {
		return 1
	}
	return num * factorial(num-1)
}

func main() {
	//fmt.Println(factorial(4))
	countDown(3)
}
