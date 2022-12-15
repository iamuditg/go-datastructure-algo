package leetcode

func countKDifference(nums []int, k int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if abs(nums[i], nums[j]) == k {
				count++
			}
		}
	}
	return count
}

func abs(a int, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}
