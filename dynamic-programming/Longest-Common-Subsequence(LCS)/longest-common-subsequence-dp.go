package main

import "fmt"

/*
Problem:  https://www.geeksforgeeks.org/longest-common-subsequence-dp-4/

Input: S1 = “AGGTAB”,
       S2 = “GXTXAYB”
Output: 4

time complexity:  O(n * m))
space complexity: O(n * m)
*/

var dp [][]int

func main() {

	str1 := "XYZW"
	str2 := "XYWZ"
	n := len(str1)
	m := len(str2)

	count := longestCommonSubSequenceDP(str1, str2, n, m)
	fmt.Println(count)
}

func longestCommonSubSequenceDP(str1, str2 string, n, m int) int {
	dp = make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				s1 := dp[i-1][j]
				s2 := dp[i][j-1]
				if s1 >= s2 {
					dp[i][j] = s1
				} else {
					dp[i][j] = s2
				}
			}
		}
	}
	for i := 0; i < n+1; i++ {
		for j := 0; j < m+1; j++ {
			fmt.Print(dp[i][j], ", ")
		}
		fmt.Println()
	}
	i := n
	j := m
	str := ""
	for i > 0 && j > 0 {
		if dp[i-1][j] == dp[i][j-1] {
			if dp[i-1][j] != dp[i][j] {
				str = string(str2[j-1]) + str
				i--
				j--
			} else {
				i--
			}

		} else if dp[i-1][j] >= dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	fmt.Println(str)
	return dp[n][m]
}
