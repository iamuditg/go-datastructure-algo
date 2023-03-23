package leetcode

func threeSum(nums []int) [][]int {
	// Get the length of the input
	n := len(nums)

	// loop through all the elements and bubble sort the array in non-decreasing order
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}

	// Initialize an empty slices to store the result
	var result [][]int

	// loop through all possible values for the first element
	for i := 0; i < n-2; i++ {

		// Skip duplicates to avoid duplicate triplets
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		// Set pointers for the second and third elements
		j := i + 1
		k := n - 1

		// loop through all possible pairs of second and third elements
		// using the two-pointer approach
		for j < k {

			// calculate the sum of the three elements
			sum := nums[i] + nums[j] + nums[k]

			if sum == 0 {
				// If the sum is zero, we found a valid triplet
				result = append(result, []int{nums[i], nums[j], nums[k]})

				// skip duplicates for the second and third elements
				for j < k && nums[j] == nums[j+1] {
					j++
				}
				for j < k && nums[k] == nums[k-1] {
					k--
				}

				// move the pointers to the next distinct element
				j++
				k--
			} else if sum < 0 {
				// If the sum is negative, we need a larger sum,
				// so we move the second pointer to the right
				j++
			} else {
				// If the sum is positive, we need a smaller sum
				// so we move the third pointer to the left
				k--
			}
		}
	}

	return result
}
