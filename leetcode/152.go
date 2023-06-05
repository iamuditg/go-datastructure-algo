package leetcode

import "fmt"

func maxProduct(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	// Initialize the variables
	maxProd := nums[0]
	minProd := nums[0]
	result := nums[0]

	for i := 1; i < n; i++ {
		// if the current number is negative. swap min and
		// max products
		fmt.Println("first:", maxProd, minProd)
		if nums[i] < 0 {
			maxProd, minProd = minProd, maxProd
		}
		fmt.Println("second:", maxProd, minProd)

		// Update the max and min products
		maxProd = max(nums[i], maxProd*nums[i])
		minProd = min(nums[i], minProd*nums[i])
		fmt.Println("third:", maxProd, minProd)
		result = max(result, maxProd)
		fmt.Println("result---->", result)
	}
	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
