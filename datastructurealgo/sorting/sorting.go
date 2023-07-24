package main

import (
	"fmt"
)

// MergeSort is a sorting algorithm that uses the divide-and-conquer strategy to sort an array recursively.
func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])

	return merge(left, right)
}

// merge is a helper function to merge two sorted arrays into a single sorted array.
func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	l, r := 0, 0

	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}

	result = append(result, left[l:]...)
	result = append(result, right[r:]...)

	return result
}

// QuickSort is a sorting algorithm that uses the divide-and-conquer strategy to sort an array recursively.
func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[0]
	var left, right []int

	for _, num := range arr[1:] {
		if num <= pivot {
			left = append(left, num)
		} else {
			right = append(right, num)
		}
	}

	left = QuickSort(left)
	right = QuickSort(right)

	return append(append(left, pivot), right...)
}

// BubbleSort is a sorting algorithm that repeatedly steps through the list to be sorted, compares each pair of adjacent items, and swaps them if they are in the wrong order.
func BubbleSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	for i := 0; i < n-1; i++ {
		if arr[i] > arr[i+1] {
			arr[i], arr[i+1] = arr[i+1], arr[i]
		}
	}

	BubbleSort(arr[:n-1])
}

// SelectionSort is a sorting algorithm that divides the input list into two parts: the sublist of items already sorted and the sublist of items remaining to be sorted.
func SelectionSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	minIndex := 0
	for i := 1; i < n; i++ {
		if arr[i] < arr[minIndex] {
			minIndex = i
		}
	}

	arr[0], arr[minIndex] = arr[minIndex], arr[0]
	SelectionSort(arr[1:])
}

// InsertionSort is a sorting algorithm that builds a final sorted array one item at a time.
func InsertionSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	InsertionSort(arr[:n-1])

	last := arr[n-1]
	j := n - 2
	for j >= 0 && arr[j] > last {
		arr[j+1] = arr[j]
		j--
	}

	arr[j+1] = last
}

// ShellSort is a sorting algorithm that improves upon the insertion sort algorithm by breaking the original list into smaller sublists.
func ShellSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	gap := n / 2
	for gap > 0 {
		for i := gap; i < n; i++ {
			temp := arr[i]
			j := i

			for j >= gap && arr[j-gap] > temp {
				arr[j] = arr[j-gap]
				j -= gap
			}

			arr[j] = temp
		}

		gap /= 2
	}
}

// HeapSort is a sorting algorithm that uses a binary heap data structure to create a partially sorted array.
func HeapSort(arr []int) {
	n := len(arr)
	if n <= 1 {
		return
	}

	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0)
	}
}

func heapify(arr []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left] > arr[largest] {
		largest = left
	}

	if right < n && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}

// --------------------------------------------------------------
// Test and Benchmark functions

// Test sorting algorithms

// testSortingAlgorithm is a helper function to test a sorting algorithm on a given array.
func testSortingAlgorithm(sortFunc func([]int), algorithmName string, arr []int) {
	fmt.Printf("Testing %s:\n", algorithmName)
	fmt.Printf("Original array: %v\n", arr)

	sortFunc(arr)

	fmt.Printf("Sorted array: %v\n", arr)
	fmt.Println("---------------------------")
}

// Benchmark sorting algorithms

func benchmarkSortingAlgorithm(sortFunc func([]int), arr []int) {
	for i := 0; i < 100000; i++ {
		sortFunc(arr)
	}
}

// Main function

func main() {
	// Sample array to test sorting algorithms
	arr := []int{64, 34, 25, 12, 22, 11, 90}

	// Test Merge Sort
	fmt.Println("Original array:", arr)
	sortedArr := MergeSort(arr)
	fmt.Println("Merge Sorted array:", sortedArr)
	fmt.Println("---------------------------")

	// Test Quick Sort
	fmt.Println("Original array:", arr)
	sortedArr = QuickSort(arr)
	fmt.Println("Quick Sorted array:", sortedArr)
	fmt.Println("---------------------------")

}
