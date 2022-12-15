package leetcode

func countGoodTriplets(arr []int, a int, b int, c int) int {
	count := 0
	for i := 0; i < len(arr)-2; i++ {
		for j := i + 1; j < len(arr)-1; j++ {
			for k := j + 1; k < len(arr); k++ {
				if absVal(arr[i]-arr[j]) <= a && absVal(arr[j]-arr[k]) <= b && absVal(arr[i]-arr[k]) <= c {
					count++
				}
			}
		}
	}
	return count
}
func absVal(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}
