package main

import "fmt"

/*
Input: arr[] = {1, 2, 3, 3}, X = 6
Output: 3

time complexity O(2^n)
space complexity O(n)

*/
func main() {
	arr := []int{1, 2, 3, 3}
	size := 4
	sum := 6
	count := countSubsetsSum(arr, size, sum)
	fmt.Println(count)
}

func countSubsetsSum(arr []int, size, sum int) int {
	if size == 0 && sum != 0 {
		return 0
	}
	if sum == 0 {
		return 1
	}

	if arr[size-1] <= sum {
		return countSubsetsSum(arr, size-1, sum-arr[size-1]) + countSubsetsSum(arr, size-1, sum)
	}
	return countSubsetsSum(arr, size-1, sum)
}
