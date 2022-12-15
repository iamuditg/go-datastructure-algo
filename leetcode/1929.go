package leetcode

//1929. Concatenation of Array

func GetConcatenation(nums []int) []int {
	output := append(nums, nums...)
	return output
}
