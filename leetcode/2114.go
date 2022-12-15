package leetcode

// 2114. Maximum Number of Words Found in Sentences

func MostWordsFound(sentences []string) int {
	output, max := 0, 0
	for i := 0; i < len(sentences); i++ {
		for j := 0; j < len(sentences[i]); j++ {
			if sentences[i] == " " {
				output++
			}
		}
		if output > max {
			max = output
		}
		output = 0
	}
	return output
}
