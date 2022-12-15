package leetcode

func maxProductDifference(nums []int) int {
	// sort array value
	for i := 0; i < len(nums); i++ {
		for j := 1 + i; j < len(nums); j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

	return (nums[len(nums)-1] * nums[len(nums)-2]) - (nums[0] * nums[1])
}
