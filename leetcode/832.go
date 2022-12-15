package leetcode

func flipAndInvertImage(image [][]int) [][]int {
	var result [][]int
	var val []int
	for i := 0; i < len(image); i++ {
		for j := len(image) - 1; j >= 0; j-- {
			element := invert(image[i][j])
			val = append(val, element)
		}
		result = append(result, val)
		val = []int{}
	}
	return result
}

func invert(i int) int {
	if i == 0 {
		return 1
	}
	return 0
}
