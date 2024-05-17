package main

import "fmt"

/*
https://www.geeksforgeeks.org/number-of-ways-to-calculate-a-target-number-using-only-array-elements/

Input : N = 5, arr[] = {1, 1, 1, 1, 1}, target = 3
Output: 5


time complexity :   O(2^n)
space complexity :   O(size of array)

*/

func main() {

	arr := []int{1, 1, 1, 1, 1}
	size := len(arr)
	target := 3 // targte or diff
	sum := arraySum(arr, target)

	count := countSubsetSumRecursion(arr, size, sum)
	fmt.Println(count)
}

func countSubsetSumRecursion(arr []int, size, sum int) int {

	if size == 0 && sum != 0 {
		return 0
	}
	if sum == 0 {
		return 1
	}

	count := 0
	if arr[size-1] <= sum {
		count = countSubsetSumRecursion(arr, size-1, sum-arr[size-1]) + countSubsetSumRecursion(arr, size-1, sum)
	} else {
		count = countSubsetSumRecursion(arr, size-1, sum)
	}

	return count
}

func arraySum(arr []int, diff int) int {

	// formula : s1 = (arraysum-diff)/2
	// s1 is sum of subset
	sum := 0
	for _, v := range arr {
		sum += v
	}
	tagetSum := (sum - diff) / 2
	return tagetSum
}
