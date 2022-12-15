package leetcode

// 1365. How Many Numbers Are Smaller Than the Current Number

func smallerNumbersThanCurrent(nums []int) []int {
	output := make([]int, len(nums))
	count := 0
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums); j++ {
			if nums[j] < nums[i] {
				count++
			}
		}
		output[i] = count
		count = 0
	}
	return output
}
