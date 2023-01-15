package leetcode

func singleNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dic := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dic[nums[i]] += 1
	}

	for i := 0; i < len(nums); i++ {
		if dic[nums[i]] == 1 {
			return nums[i]
		}
	}
	return 0
}
