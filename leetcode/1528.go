package leetcode

// 1528. Shuffle String

func restoreString(s string, indices []int) string {
	output := make([]byte, len(indices))
	for i := 0; i < len(s); i++ {
		output[indices[i]] = s[i]
	}
	return string(output)
}
