package leetcode

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandies := 0
	// find max value in candies
	for i := 0; i < len(candies); i++ {
		if maxCandies < candies[i] {
			maxCandies = candies[i]
		}
	}
	// find bool value of candies
	boolCandies := make([]bool, len(candies))
	for j := 0; j < len(candies); j++ {
		if candies[j]+extraCandies >= maxCandies {
			boolCandies[j] = true
		} else {
			boolCandies[j] = false
		}
	}
	return boolCandies
}
