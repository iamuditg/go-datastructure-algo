package leetcode

func sumOddLengthSubarrays(arr []int) int {
	sum := 0
	for a := 0; a < len(arr); a++ {
		for b := a; b < len(arr); b += 2 {
			for c := a; c <= b; c++ {
				sum += arr[c]
			}
		}
	}
	return sum
}
