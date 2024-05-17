package main

import (
	"fmt"
)

/*
Input: sum = 4, coins[] = {1,2,3},
Output: 4
Explanation: there are four solutions: {1, 1, 1, 1}, {1, 1, 2}, {2, 2}, {1, 3}.

time complexity:   O(size of array * sum)
space complexity: O(size of array * sum)
*/

func main() {

	coin := []int{1, 2, 3}
	size := len(coin)
	sum := 4
	count := maxNumOfWayForCoinChange(coin, size, sum)
	fmt.Println(count)
}

func maxNumOfWayForCoinChange(coin []int, size, sum int) int {
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
			if coin[i-1] <= j {
				dp[i][j] = dp[i][j-coin[i-1]] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	return dp[size][sum]
}
