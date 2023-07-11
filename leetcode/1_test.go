package leetcode

import "testing"

// PROBLEM - https://leetcode.com/problems/two-sum/

// twoSumOne finds two numbers in the given array that add up to the target
// It returns the indices of the two numbers as a slice.
func twoSumOne(nums []int, target int) []int {
	var output []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				output = append(output, i, j)
				break
			}
		}
	}
	return output
}

// twoSum finds two numbers in the given array that add up to the target
// It returns the indices of the two numbers as a slice.
func twoSum(nums []int, target int) []int {
	var output []int
	numsMap := make(map[int]int)

	// Build a map of number to index
	for key, value := range nums {
		numsMap[value] = key
	}

	// Check if the complement exists in the map
	for key, value := range nums {
		result := target - value
		if nextKey, exist := numsMap[result]; exist && nextKey != key {
			output = append(output, key, nextKey)
			break
		}
	}

	return output
}

// twoSum1 finds two numbers in the given array that add up to the target.
// It returns the indices of the two numbers as a slice.
func twoSum1(nums []int, target int) []int {
	dict := make(map[int]int)

	for pos, num := range nums {
		pairValue := target - num
		pairPos, ok := dict[pairValue]
		if ok && pairPos != pos {
			return []int{pairPos, pos}
		}
		dict[num] = pos
	}
	return []int{-1, -1}
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
	}

	for _, test := range tests {
		result := twoSum(test.nums, test.target)
		if len(result) != len(test.expected) {
			t.Errorf("twoSum(%v, %d) = %v, expected %v", test.nums, test.target, result, test.expected)
		} else {
			for i := range result {
				if result[i] != test.expected[i] {
					t.Errorf("twoSum(%v, %d) = %v, expected %v", test.nums, test.target, result, test.expected)
					break
				}
			}
		}
	}
}

func BenchmarkTwoSum(b *testing.B) {
	nums := []int{2, 7, 11, 15}
	target := 9
	for i := 0; i < b.N; i++ {
		twoSum(nums, target)
	}
}
