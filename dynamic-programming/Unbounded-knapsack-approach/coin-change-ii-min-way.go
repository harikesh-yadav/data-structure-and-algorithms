package main

import (
	"fmt"
	"math"
)

/*
https://www.geeksforgeeks.org/find-minimum-number-of-coins-that-make-a-change/

video:  https://www.youtube.com/watch?v=I-l6PBeERuc&list=PL_z_8CaSLPWekqhdCPmFohncHwz8TY2Go&index=17

Input: coins[] = {9, 6, 5, 1}, V = 11
Output: Minimum 2 coins required
We can use one coin of 6 cents and 1 coin of 5 cents


time complexity :    O(size * sum)
space complexity: O(size * sum)
*/

var dp [][]int

func main() {
	coins := []int{9, 6, 5, 1}
	size := len(coins)
	sum := 11

	mincount := minNumberOfWayForCoinChange(coins, size, sum)
	fmt.Println(mincount)
}

func minNumberOfWayForCoinChange(coins []int, size, sum int) int {

	dp = make([][]int, size+1)
	for i := 0; i < size+1; i++ {
		dp[i] = make([]int, sum+1)
	}

	for i := 0; i < size+1; i++ {
		for j := 0; j < sum+1; j++ {
			if i == 0 {
				dp[i][j] = math.MaxInt - 1
			}
			if j == 0 {
				dp[i][j] = 0
			}
			if i == 1 {
				if j%coins[i-1] == 0 {
					dp[i][j] = j
				} else {
					dp[i][j] = math.MaxInt - 1
				}
			}
		}
	}

	for i := 2; i < size+1; i++ {
		for j := 0; j < sum+1; j++ {
			if coins[i-1] <= j {
				s1 := dp[i][j-coins[i-1]] + 1
				s2 := dp[i-1][j]
				if s1 < s2 {
					dp[i][j] = s1
				} else {
					dp[i][j] = s2
				}
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[size][sum]
}
