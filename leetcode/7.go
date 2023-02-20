package leetcode

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	ans := 0
	for x != 0 {
		last := x % 10
		x /= 10
		next := ans*10 + last
		fmt.Println("last: ", last, " x: ", x, " next: ", next, " (next-last)/10 ", (next-last)/10, " ans: ", ans)
		if (next-last)/10 != ans || next > math.MaxInt32 || next < math.MinInt32 {
			return 0
		}
		ans = next
		fmt.Println("-------")
	}
	return ans
}
