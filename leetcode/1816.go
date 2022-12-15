package leetcode

func truncateSentence(s string, k int) string {
	byteStr := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if string(s[i]) == " " {
			k--
			if k == 0 {
				break
			}
		}
		byteStr = append(byteStr, s[i])
	}
	return string(byteStr)
}
