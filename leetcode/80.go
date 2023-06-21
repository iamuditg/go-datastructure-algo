package leetcode

func RemoveDuplicates(nums []int) int {
	// If the array is empty, return 0
	if len(nums) == 0 {
		return 0
	}

	// Initialize variables
	count := 1      // Number of occurrences of the current element
	duplicates := 0 // Number of duplicates encountered so far
	prev := nums[0] // Previous element

	for i := 1; i < len(nums); i++ {
		if nums[i] == prev {
			// If the current element is equal to the previous element
			count++
			if count > 2 {
				// If there are more than 2 occurrences, it's a duplicate
				duplicates++
			}
		} else {
			// If the current element is different from the previous element
			count = 1      // Reset the count
			prev = nums[i] // Update the previous element
		}

		// Move the element forward to overwrite duplicates if necessary
		nums[i-duplicates] = nums[i]
	}

	// Return the new length of the modified array
	return len(nums) - duplicates
}
