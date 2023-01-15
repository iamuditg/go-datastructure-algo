package leetcode

func majorityElement(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	dic := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dic[nums[i]] += 1
	}

	count := 0
	result := 0
	for k, v := range dic {
		if count < v {
			count = v
			result = k
		}
	}
	return result
}
