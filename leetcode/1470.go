package leetcode

// 1470. Shuffle the Array

func shuffle(nums []int, n int) []int {
	output := make([]int, len(nums))
	l := 0
	r := n

	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			output[i] = nums[l]
			l++
		} else {
			output[i] = nums[r]
			r++
		}
	}
	return output
}
