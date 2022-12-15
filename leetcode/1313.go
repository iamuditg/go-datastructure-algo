package leetcode

// 1313. Decompress Run-Length Encoded List

func decompressRLElist(nums []int) []int {
	output := make([]int, 0)

	for i := 0; i < len(nums); i += 2 {
		freq := nums[i+1]
		val := nums[i]
		for j := val; j > 0; j-- {
			output = append(output, freq)
		}
	}
	return output
}
