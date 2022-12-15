package leetcode

func numIdenticalPairs(nums []int) int {
	pairCount := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] && i < j {
				pairCount++
			}
		}
	}
	return pairCount
}
