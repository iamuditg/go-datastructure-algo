package main

import "fmt"

/*
 Given an array of integer, find the maximum sum of any contiguous subarray.
*/

func MaxSubArraySum(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	maxSum := arr[0]
	currentSum := arr[0]

	for i := 1; i < len(arr); i++ {
		fmt.Println(currentSum, arr[i], maxSum)
		if currentSum+arr[i] > arr[i] {
			currentSum += arr[i]
		} else {
			currentSum = arr[i]
		}

		if currentSum > maxSum {
			maxSum = currentSum
		}
	}
	return maxSum
}
