package leetcode

func countConsistentStrings(allowed string, words []string) int {
	count := 0
	for _, word := range words {
		isValid := true
		for _, r1 := range word {
			has := false
			for _, r2 := range allowed {
				if r1 == r2 {
					has = true
					break
				}
			}
			if !has {
				isValid = false
				break
			}
		}
		if isValid {
			count++
		}
	}
	return count
}
