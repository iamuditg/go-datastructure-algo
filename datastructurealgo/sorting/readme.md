## Sorting Algorithms using Recursion

This repository contains various sorting algorithms implemented using recursion in Go.

## Sorting Algorithms Implemented

### 1. Merge Sort

**Method Signature:** `func MergeSort(arr []int) []int`

Merge Sort is a divide-and-conquer algorithm that recursively divides the input array into two halves, sorts them separately, and then merges them to produce the final sorted output. It has a time complexity of O(n log n) and is considered one of the most efficient sorting algorithms.

### 2. Quick Sort

**Method Signature:** `func QuickSort(arr []int) []int`

Quick Sort is another divide-and-conquer algorithm that selects a pivot element, partitions the array around the pivot, and recursively sorts the two resulting subarrays. It has an average time complexity of O(n log n), making it one of the fastest sorting algorithms in practice.

### 3. Bubble Sort

**Method Signature:** `func BubbleSort(arr []int)`

Bubble Sort repeatedly compares adjacent elements in the array and swaps them if they are in the wrong order. This process is repeated until the entire array is sorted. It has a time complexity of O(n^2) and is not recommended for large datasets due to its inefficiency.

### 4. Selection Sort

**Method Signature:** `func SelectionSort(arr []int)`

Selection Sort divides the array into a sorted and an unsorted portion, iteratively finds the minimum element from the unsorted portion, and places it at the end of the sorted portion. It has a time complexity of O(n^2) but performs better than Bubble Sort in practice.

### 5. Insertion Sort

**Method Signature:** `func InsertionSort(arr []int)`

Insertion Sort builds a sorted array one element at a time by repeatedly inserting elements from the unsorted part into the correct position in the sorted part. It has a time complexity of O(n^2) but performs well on small datasets or partially sorted arrays.

### 6. Shell Sort

**Method Signature:** `func ShellSort(arr []int)`

Shell Sort is an extension of insertion sort where the array is sorted in "h-sorting" with decreasing values of h. It has a time complexity that depends on the gap sequence used, but in the worst case, it can be O(n^2). It performs better than insertion sort on average.

### 7. Heap Sort

**Method Signature:** `func HeapSort(arr []int)`

Heap Sort creates a binary heap data structure to build a partially sorted array and then repeatedly extracts the maximum element to obtain the sorted array. It has a time complexity of O(n log n) and is an in-place sorting algorithm.

## Test and Benchmark

The `TestSortingAlgorithms` function is provided to test each sorting algorithm with sample inputs, and the `BenchmarkSortingAlgorithms` function is used to benchmark the performance of the sorting algorithms.

Please refer to the respective implementations in the code for detailed explanations and method definitions.