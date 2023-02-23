package leetcode

import "fmt"

// Function to generate all combination of well-formed parentheses
func generateParenthesis(n int) []string {
	// Initialize an empty slice to store the result
	result := []string{}

	// call the backtrack function to generate all combination of parentheses
	backtrack(&result, "", 0, 0, n)

	// return the result
	return result
}

// backtrack function to generate all combination of well-formed parenthesis
func backtrack(result *[]string, current string, open int, close int, max int) {
	// if the length of the current string is equal to twice the input n, it means we have
	//generated a well-formed combination of parentheses
	fmt.Println(current)
	if len(current) == max*2 {
		// append the well-formed combination of parentheses to the result slice
		*result = append(*result, current)
		// return from the function
		return
	}

	// if we have not used up all the opening brackets, we can add an opening brackets to the current string
	if open < max {
		backtrack(result, current+"(", open+1, close, max)
	}

	// if we have more closing brackets than opening brackets, we can add a closing bracket to the current string
	if close < open {
		backtrack(result, current+")", open, close+1, max)
	}

}
