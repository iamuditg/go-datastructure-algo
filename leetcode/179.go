package leetcode

func largestNumber(nums []int) string {

	// convert int to string
	strNum := make([]string, len(nums))
	for i, num := range nums {
		strNum[i] = intToString(num)
	}

	// sort the strNum is descending order
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-i-1; j++ {
			if compare(strNum[j], strNum[j+1]) == -1 {
				strNum[j], strNum[j+1] = strNum[j+1], strNum[j]
			}
		}
	}

	// append string to form largest Number
	var result string
	for _, str := range strNum {
		result += str
	}

	// remove leading zeroes from result
	result = removeLeadingZeroes(result)

	return result
}

func intToString(num int) string {

	// if num is zero return zero string
	if num == 0 {
		return "0"
	}

	// check if num is negative or positive
	isNegative := false
	if num < 0 {
		isNegative = true
		num = -num
	}

	var strNum []byte
	for num > 0 {
		digit := num % 10
		strNum = append(strNum, byte(digit)+'0')
		num /= 10
	}

	if isNegative {
		strNum = append(strNum, '-')
	}

	reverseNum(strNum)

	return string(strNum)
}

func reverseNum(num []byte) {
	i, j := 0, len(num)-1
	for i < j {
		num[i], num[j] = num[j], num[i]
		i++
		j--
	}
}

func compare(a, b string) int {
	ab := a + b
	ba := b + a

	if ab > ba {
		return 1
	} else if ab < ba {
		return -1
	}

	return 0
}

func removeLeadingZeroes(str string) string {
	if str == "" {
		return str
	}

	i := 0
	for i < len(str)-1 && str[i] == '0' {
		i++
	}

	return str[i:]
}
