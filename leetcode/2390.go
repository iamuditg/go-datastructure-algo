package leetcode

func removeStars(s string) string {
	str := []byte{}

	for i := 0; i < len(s); i++ {
		if s[i] == '*' {
			if len(str) > 0 {
				str = str[:len(str)-1]
			}
		} else {
			str = append(str, s[i])
		}
	}
	return string(str)
}