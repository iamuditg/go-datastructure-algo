package leetcode

func searchRange(nums []int, target int) []int {
	first, second := -1, -1
	for i, value := range nums {
		if value == target {
			if first == -1 {
				first = i
				second = i
			} else {
				second = i
			}
		}
	}

	return []int{first, second}
}
