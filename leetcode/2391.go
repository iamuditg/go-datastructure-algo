package leetcode

import (
	"fmt"
	"strings"
)

func garbageCollectionOne(garbage []string, travel []int) int {
	travel = append(travel[:1], travel...)
	travel[0] = 0
	fmt.Println(travel)
	trucks := []string{"G", "P", "M"}
	count := 0
	for _, s := range trucks {
		res := 0
		idx := 0
		for i := 0; i < len(garbage); i++ {
			c := strings.Count(garbage[i], s)
			fmt.Println(c, garbage[i], s)
			if c > 0 {
				res += c
				idx = i
			}
			fmt.Println("res: ", res, "idx: ", idx)
		}
		sum := 0
		for j := 0; j < idx+1; j++ {
			sum += travel[j]
		}
		count += res + sum
	}
	return count
}
