package leetcode

import "fmt"

func wordBreak(s string, wordDict []string) bool {

	// create a set for faster word lookup
	wordSet := make(map[string]bool)

	for _, word := range wordDict {
		wordSet[word] = true
	}

	fmt.Println(wordSet)

	n := len(s)
	dp := make([]bool, n+1)
	dp[0] = true

	fmt.Println(dp)
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			// check if the substring s[j:i] is in the word dictionary and
			// if the previous substring s[0:j] can be segmented
			fmt.Println("i: ", i, " j: ", j, " dp[j]: ", dp[j], " s[j:i]: ", s[j:i], " wordSet[s[j:i]]: ", wordSet[s[j:i]], " actualWord: ", s)
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				fmt.Println("inside: ", dp)
				break
			}
		}
	}

	fmt.Println("dp[n] : ", dp[n])
	return dp[n]
}
