package main

func checkArraySorted(a []int, n int) bool {
	if n == 0 || n == 1 {
		return true
	}
	if a[0] > a[1] {
		return false
	}
	isSmallSorted := checkArraySorted(a[1:], n-1)
	return isSmallSorted
}

func sumOfArray(a []int, n int) int {
	if n == 0 {
		return 0
	}
	sum := a[0] + sumOfArray(a[1:], n-1)
	return sum
}

func checkIfPresent(a []int, n int, element int) bool {
	if n == 0 {
		return false
	}
	if a[0] == element {
		return true
	}
	check := checkIfPresent(a[1:], n-1, element)
	return check
}

func firstIndexOfElement(a []int, n int, i int, element int) int {
	if i == n {
		return -1
	}
	if a[i] == element {
		return i
	}
	index := firstIndexOfElement(a, n, i+1, element)
	return index
}

func lastIndexOfElement(a []int, n int, i int, element int) int {
	if n == i {
		return -1
	}
	if a[n] == element {
		return n
	}
	index := lastIndexOfElement(a, n-1, i, element)
	return index
}

func printAllPosition(a []int, n int, element int) {
	if n < 0 {
		return
	}
	if a[n] == element {
		println(n)
	}
	printAllPosition(a, n-1, element)

}

func storeAllPosition(a []int, n int, element int, store []int) []int {
	if n < 0 {
		return store
	}
	if a[n] == element {
		store = append(store, n)
	}
	store = storeAllPosition(a, n-1, element, store)
	return store
}

func checkPalindrome(str string, i int, n int) bool {
	if i > n {
		return true
	}
	if str[i] != str[n] {
		return false
	}
	smallAns := checkPalindrome(str, i+1, n-1)
	return smallAns
}
