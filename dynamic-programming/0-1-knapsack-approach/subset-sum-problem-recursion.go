package main

import (
	"fmt"
)

/*
Input: set[] = {3, 34, 4, 12, 5, 2}, sum = 9
Output: True


time complexity O(2^n)
space complexity   O(n)

*/

func main() {
	arr := []int{3, 34, 4, 12, 5, 2}
	sum := 9
	size := len(arr)
	isExist := subsetSumRecursion(arr, size, sum)
	fmt.Println(isExist)
}

func subsetSumRecursion(arr []int, size, sum int) bool {

	if size == 0 && sum != 0 {
		return false
	}
	if sum == 0 {
		return true
	}

	if arr[size-1] <= sum {
		return subsetSumRecursion(arr, size-1, sum-arr[size-1]) || subsetSumRecursion(arr, size-1, sum)
	}
	return subsetSumRecursion(arr, size-1, sum)

}
