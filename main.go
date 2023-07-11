// CODE EXAMPLE VALID FOR COMPILING
package main

import (
	"fmt"
)

/*

Problem- Write a Program in golang
1. n different goroutine insert any random value into array/slice i.e insert i= 0..10 in arr []int.
and after insert print the length of array in main().
*/

/*
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
func twoSum(nums []int, target int) []int {
}
*/

func twoSum(nums []int, target int) []int {
	output := []int{-1, -1}

	visited := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if v, exists := visited[target-nums[i]]; exists {
			output[0] = v
			output[1] = i
		}
		visited[nums[i]] = i
	}
	return output
}

/*
	type Str struct{
	    array []int
	    mx sync.Mutex
	}
*/
func main() {
	/*fmt.Println("hello world")
	  n:=10
	  obj:= Str{array:make([]int,10)}
	  var wg sync.WaitGroup
	  for i:=0;i<n;i++{
	      wg.Add(1)
	      i:=i
	      go func(i int,obj Str,wg *sync.WaitGroup){
	          defer wg.Done()
	              obj.mx.Lock()
	              obj.array[i]=i
	              //obj.counter++
	              obj.mx.Unlock()
	          }(i,obj,&wg)
	  }
	  wg.Wait()

	  fmt.Printf("Array: %v, Len: %d",obj.array,len(obj.array))
	  //fmt.Println("End")*/
	fmt.Printf("Array: %v", twoSum([]int{2, 7, 11, 15}, 18))
}
