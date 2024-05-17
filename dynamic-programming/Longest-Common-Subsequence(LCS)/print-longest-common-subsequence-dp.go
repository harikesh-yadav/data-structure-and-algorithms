package main

import "fmt"

/**
https://www.geeksforgeeks.org/printing-longest-common-subsequence/

Input: S1 = "AGGTAB", S2 = "GXTXAYB"
Output: 4


time complexity:  O(n * m)

space complexity:  O(n * m)

**/

var dp [][]int

func main() {
	str1 := "AGGTAB"
	str2 := "GXTXAYB"
	n := len(str1)
	m := len(str2)

	s := printLongestCommonSubSequence(str1, str2, n, m)
	fmt.Println(s)
}

func printLongestCommonSubSequence(str1, str2 string, n, m int) string {

	dp = make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
	}

	for i := 0; i < n+1; i++ {
		for j := 0; j < m+1; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
			}
		}
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
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

	i := n
	j := m
	str := ""

	for i > 0 && j > 0 {
		if str1[i-1] == str2[j-1] {
			str = string(str1[i-1]) + str
			i--
			j--
		} else {
			if dp[i][j-1] > dp[i-1][j] {
				j--
			} else {
				i--
			}
		}
	}

	return str

}
