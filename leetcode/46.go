package leetcode

func permute(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{}
	}
	var result [][]int
	visited := make([]bool, len(nums))
	backtrackPermute(&result, visited, []int{}, nums)
	return result
}

func backtrackPermute(result *[][]int, visited []bool, current, nums []int) {
	// base condition
	if len(current) == len(nums) {
		*result = append(*result, append([]int{}, current...))
		return
	}

	for i := 0; i < len(nums); i++ {
		if visited[i] == true {
			continue
		}
		visited[i] = true
		current = append(current, nums[i])
		backtrackPermute(result, visited, current, nums)
		visited[i] = false
		current = current[:len(current)-1]
	}
}
