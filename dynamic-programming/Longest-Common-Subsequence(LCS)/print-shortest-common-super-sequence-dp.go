package main

import "fmt"

/***
problem: https://www.geeksforgeeks.org/print-shortest-common-supersequence/


Input:Input: X = "AGGTAB",  Y = "GXTXAYB"
Output: Output: "AGXGTXAYB" OR "AGGXTXAYB"

time complexity:  O(n * m))
space complexity: O(n * m)

**/

var dp [][]int

func main() {

	str1 := "AGGTAB"
	str2 := "GXTXAYB"

	n := len(str1)
	m := len(str2)

	str := printShortestCommonSuperSequence(str1, str2, n, m)

	fmt.Println(str)

}

func printShortestCommonSuperSequence(str1, str2 string, n, m int) string {
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

	i := n
	j := m
	str := ""
	for i > 0 && j > 0 {
		if str1[i-1] == str2[j-1] {
			str = string(str1[i-1]) + str
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			str = string(str2[j-1]) + str
			j--
		} else {
			str = string(str1[i-1]) + str
			i--
		}
	}

	for i > 0 {
		str = string(str1[i-1]) + str
		i--
	}
	for j > 0 {
		str = string(str2[j-1]) + str
		j--
	}
	return str

}
