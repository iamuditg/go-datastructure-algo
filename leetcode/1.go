package leetcode

import "fmt"

func twoSumOne(nums []int, target int) []int {
	var output []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				output = append(output, i, j)
				break
			}
		}
	}
	return output
}

func twoSum(nums []int, target int) []int {
	var output []int
	numsMap := make(map[int]int)
	fmt.Println(numsMap)
	for key, value := range nums {
		numsMap[value] = key
	}
	fmt.Println(numsMap)

	for key, value := range nums {
		result := target - value
		if nextKey, exist := numsMap[result]; exist && nextKey != key {
			fmt.Println(nextKey, key)
			output = append(output, key, nextKey)
			break
		}
	}
	return output
}

func twoSum1(nums []int, target int) []int {
	dict := make(map[int]int)
	for pos, num := range nums {
		pairValue := target - num
		pairPos, ok := dict[pairValue]
		if ok && pairPos != pos {
			return []int{pairPos, pos}
		}
		dict[num] = pos
	}
	return []int{-1, -1}
}
