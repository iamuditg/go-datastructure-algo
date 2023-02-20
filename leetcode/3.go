package leetcode

func lengthOfLongestSubstring(s string) int {
	maxr := 0
	m := [128]int{}
	for i := 0; i < 128; i++ {
		m[i] = -1
	}
	lastIndex := 0
	length := 0
	for k := 0; k < len(s); k++ {
		if m[s[k]] >= 0 {
			for i := lastIndex; i < m[s[k]]; i++ {
				m[s[i]] = -1
			}
			length -= m[s[k]] - lastIndex
			lastIndex = m[s[k]] + 1
		} else {
			length += 1
		}
		m[s[k]] = k
		if maxr < length {
			maxr = length
		}
	}
	return maxr
}
