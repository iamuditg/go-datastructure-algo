package main

import (
	"reflect"
	"testing"
)

func rotateImage(matrix [][]int) {
	n := len(matrix)
	m := len(matrix[0])
	if n != m {
		return // matrix is not square, cannot rotate
	}
	for i := 0; i < n/2; i++ {
		for j := i; j < n-i-1; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[n-j-1][i]
			matrix[n-j-1][i] = matrix[n-i-1][n-j-1]
			matrix[n-i-1][n-j-1] = matrix[j][n-i-1]
			matrix[j][n-i-1] = temp
		}
	}
}

func TestRotateImage(t *testing.T) {
	testCases := []struct {
		input  [][]int
		output [][]int
	}{
		{
			input: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			output: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			input: [][]int{
				{1, 2},
				{3, 4},
			},
			output: [][]int{
				{3, 1},
				{4, 2},
			},
		},
		{
			input: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
			},
			output: [][]int{
				{5, 1},
				{6, 2},
				{7, 3},
				{8, 4},
			},
		},
		{
			input: [][]int{
				{1, 2},
			},
			output: [][]int{
				{1},
				{2},
			},
		},
		{
			input: [][]int{
				{1},
			},
			output: [][]int{
				{1},
			},
		},
	}

	for _, tc := range testCases {
		rotateImage(tc.input)
		if !reflect.DeepEqual(tc.input, tc.output) {
			t.Errorf("unexpected output: got %v, want %v", tc.input, tc.output)
		}
	}
}
