package main

import "fmt"

func singleNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dic := make(map[int]int, len(nums))
	for i := 0; i < len(dic); i++ {
		dic[nums[i]] += 1
	}
	fmt.Println(dic)
	for i := 0; i < len(nums); i++ {
		if dic[nums[i]] == 1 {
			return nums[i]
		}
	}
	return 0
}

func main() {
	fmt.Println(singleNumber([]int{1, 2, 2}))
}
