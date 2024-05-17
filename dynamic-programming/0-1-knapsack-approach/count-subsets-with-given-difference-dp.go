package main

import "fmt"

/*

Input: N = 4, arr[] = [5, 2, 6, 4], diff = 3
Output: 1

explanation: s1-s2=3  eq-1
			s1+s2 = 7  eq-2
			s1 = (diff-array(sum))/2

time complexity O(size*(sum of subset))
space complexity O(size* (array sum))
*/
var dp [][]int

func main() {
	arr := []int{1, 2, 3, 1, 2}
	diff := 1
	size := len(arr)
	sum := subsetSum(arr, diff)

	count := countSubsetsWithDifference(arr, size, sum)
	fmt.Println(count)
}

func countSubsetsWithDifference(arr []int, size, sum int) int {
	dp = make([][]int, size+1)
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

func subsetSum(arr []int, diff int) int {
	// formula ->    s1 = (array(sum) - diff)/2
	arrsum := 0
	for _, v := range arr {
		arrsum += v
	}
	sum := (arrsum - diff) / 2
	return sum
}
