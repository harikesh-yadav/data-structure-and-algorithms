package main

import "fmt"

/**
problem:  https://www.geeksforgeeks.org/longest-repeating-subsequence/

input: "AABEBCDD"
output:3

time complexity O(n * n)
space complexity O(n * n)
**/

func main() {
	str1 := "AABEBCDD"
	n := len(str1)
	count := longestRepeatingSubSequence(str1, n)
	fmt.Println(count)
}

func longestRepeatingSubSequence(str1 string, n int) int {
	str2 := str1
	dp := make([][]int, n+1)

	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, n+1)
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < n+1; j++ {
			if str1[i-1] == str2[j-1] && i != j {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				s1 := dp[i-1][j]
				s2 := dp[i][j-1]
				if s1 > s2 {
					dp[i][j] = s1
				} else {
					dp[i][j] = s2
				}
			}
		}
	}
	return dp[n][n]
}
