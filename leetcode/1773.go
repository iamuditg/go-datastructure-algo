package leetcode

// 1773. Count Items Matching a Rule

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	index := 0
	switch ruleKey {
	case "type":
		index = 0
	case "color":
		index = 1
	case "name":
		index = 2
	}
	count := 0
	for i := 0; i < len(items); i++ {
		if items[i][index] == ruleValue {
			count++
		}
	}
	return count
}
