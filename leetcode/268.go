package leetcode

import "fmt"

func missingNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dic := make([]int, len(nums)+1)
	for i := 0; i < len(nums); i++ {
		dic[nums[i]] = 1
	}
	fmt.Println(dic)
	for i := 0; i < len(dic); i++ {
		if dic[i] == 0 {
			return i
		}
	}
	return 0
}
