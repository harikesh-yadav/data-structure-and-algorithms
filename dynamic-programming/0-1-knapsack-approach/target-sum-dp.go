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

	count := countSubsetSumDP(arr, size, sum)
	fmt.Println(count)
}

func countSubsetSumDP(arr []int, size, sum int) int {

	dp := make([][]int, size+1)
	for i := 0; i < size+1; i++ {
		dp[i] = make([]int, sum+1)
	}
	for i := 0; i < size+1; i++ {
		for j := 0; j < sum+1; j++ {
			if i == 0 {
				dp[i][j] = 0
			}
			if j == 0 {
				dp[i][j] = 1
			}
		}
	}

	for i := 1; i < size+1; i++ {
		for j := 1; j < sum+1; j++ {
			if arr[i-1] <= j {
				dp[i][j] = dp[i-1][j-arr[i-1]] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[size][sum]
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
