package main

import "fmt"

func zeroMatrix(matrix [][]int) {
	m := len(matrix)
	n := len(matrix[0])

	rows := make([]bool, m)
	columns := make([]bool, n)

	// Mark the rows and columns that need to be zeroed out
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				rows[i] = true
				columns[j] = true
			}
		}
	}

	// Zero out the marked rows
	for i := 0; i < m; i++ {
		if rows[i] {
			for j := 0; j < n; j++ {
				matrix[i][j] = 0
			}
		}
	}

	// zero out the marked columns
	for j := 0; j < n; j++ {
		if columns[j] {
			for i := 0; i < m; i++ {
				matrix[i][j] = 0
			}
		}
	}
}

func printMatrix(matrix [][]int) {
	for _, row := range matrix {
		fmt.Println(row)
		fmt.Println()
	}
}
