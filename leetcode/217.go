package leetcode

func containsDuplicate(nums []int) bool {
	if len(nums) == 0 {
		return false
	}
	dic := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		dic[nums[i]] += 1
	}
	for i := 0; i < len(nums); i++ {
		if dic[nums[i]] > 1 {
			return true
		}
	}
	return false
}
