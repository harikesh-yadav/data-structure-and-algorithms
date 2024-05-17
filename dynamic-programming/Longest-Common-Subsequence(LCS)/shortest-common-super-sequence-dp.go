package main

import "fmt"

/***
problem: https://www.geeksforgeeks.org/shortest-common-supersequence/


Input: S1 = go, S2 = “GXTXAYB”
Output: 4

time complexity:  O(n * m))
space complexity: O(n * m)

**/

var dp [][]int

func main() {

	str1 := "AGGTAB"
	str2 := "GXTXAYB"

	n := len(str1)
	m := len(str2)

	count := printShortestSuperSeq(str1, str2, n, m)

	fmt.Println(count)

}

func printShortestSuperSeq(str1, str2 string, n, m int) int {
	dp = make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
	}
	for i := 0; i < n+1; i++ {
		for j := 0; j < m+1; j++ {
			if i == 0 {
				dp[i][j] = j
			}
			if j == 0 {
				dp[i][j] = i
			}
		}
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				s1 := dp[i-1][j]
				s2 := dp[i][j-1]
				if s1 < s2 {
					dp[i][j] = s1 + 1
				} else {
					dp[i][j] = s2 + 1
				}
			}
		}
	}

	return dp[n][m]

}
