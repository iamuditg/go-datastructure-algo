package leetcode

// 1672. Richest Customer Wealth

func maximumWealth(accounts [][]int) int {
	max, output := 0, 0
	for i := 0; i < len(accounts); i++ {
		for j := 0; j < len(accounts[i]); j++ {
			output += accounts[i][j]
		}
		if output > max {
			max = output
		}
		output = 0
	}
	return max
}
