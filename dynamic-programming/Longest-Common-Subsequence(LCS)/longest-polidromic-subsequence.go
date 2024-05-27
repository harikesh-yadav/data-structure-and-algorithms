package main

import "fmt"

/**
problem: https://www.geeksforgeeks.org/longest-palindromic-subsequence-dp-12/

Input: S = “GEEKSFORGEEKS”
Output: 5
Explanation: The longest palindromic subsequence we can get is of length 5. There are more than 1 palindromic subsequences of length 5, for example: EEKEE, EESEE, EEFEE, …etc.

time complexity O(n^2)
space complexity O(n^2)
**/

func main() {
	str1 := "BBABCBCAB"
	n := len(str1)
	count := longestPolidromicSubSequence(str1, n)
	fmt.Println(count)
}

func longestPolidromicSubSequence(str1 string, n int) int {
	str2 := reverse(str1)
	dp := make([][]int, n+1)

	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < n+1; i++ {
		for j := 0; j < n+1; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
			}
		}
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < n+1; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
			} else {
				s1 := dp[i][j-1]
				s2 := dp[i-1][j]
				if s1 > s2 {
					dp[i][j] = s1
				} else {
					dp[i][j] = s2
				}
			}
		}
	}

	for i := 0; i < n+1; i++ {
		for j := 0; j < n+1; j++ {
			fmt.Print(dp[i][j], ", ")
		}
		fmt.Println()
	}

	i := n
	j := n
	str := ""

	for i > 0 && j > 0 {
		if dp[i][j-1] == dp[i-1][j] {
			if dp[i-1][j] != dp[i][j] {
				str = string(str2[j-1]) + str
				i--
				j--
			} else {
				i--
			}
		} else if dp[i][j-1] > dp[i-1][j] {
			j--
		} else {
			i--
		}
	}
	fmt.Println(str)

	return dp[n][n]
}

func reverse(str string) (s string) {
	for _, v := range str {
		s = string(v) + s
	}
	return
}
