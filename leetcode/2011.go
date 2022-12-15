package leetcode

// 2011. Final Value of Variable After Performing Operations

func finalValueAfterOperations(operations []string) int {
	output := 0
	for i := 0; i < len(operations); i++ {
		if operations[i] == "X++" || operations[i] == "++X" {
			output++
		} else if operations[i] == "X--" || operations[i] == "--X" {
			output--
		}
	}
	return output
}
