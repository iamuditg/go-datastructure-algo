package leetcode

func climbStairs(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	step1 := 1
	step2 := 1
	result := 0
	for i := 2; i <= n; i++ {
		result = step1 + step2
		step2 = step1
		step1 = result
	}
	return result
}
