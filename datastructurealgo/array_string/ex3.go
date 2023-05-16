package main

import "fmt"

/*
	Given an array of integers, rearrange the elements in such a way that all the even
	elements appear before the odd elements. the order of even elements and the order of odd
	elements should remain the same.
*/

func RearrangeArray(arr []int) []int {
	evenIdx := 0
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr, evenIdx)
		if arr[i]%2 == 0 {
			arr[i], arr[evenIdx] = arr[evenIdx], arr[i]
			evenIdx++
		}
	}
	return arr
}
