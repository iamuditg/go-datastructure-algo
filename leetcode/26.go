package leetcode

func removeDuplicates(nums []int) int {
	l := 0
	if len(nums) == 0 {
		return 0
	}
	for i := 0; i < len(nums); i++ {
		if i == 0 || nums[i] != nums[i-1] {
			nums[l] = nums[i]
			l++
		}
	}
	return l
}
