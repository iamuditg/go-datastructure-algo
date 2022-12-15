package leetcode

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	strByte := make([]byte, 0)
	strByte2 := make([]byte, 0)
	for _, val := range word1 {
		for i := 0; i < len(val); i++ {
			strByte = append(strByte, val[i])
		}
	}

	for _, val := range word2 {
		for i := 0; i < len(val); i++ {
			strByte2 = append(strByte2, val[i])
		}
	}
	return string(strByte) == string(strByte2)
}
