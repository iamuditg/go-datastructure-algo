package leetcode

func mySqrt(x int) int {
	if x == 0 {
		return x
	}
	left := 0
	right := x
	for left <= right {
		mid := (left + right) / 2
		mul := mid * mid
		if mul == x {
			return mid
		} else if mul < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right
}
