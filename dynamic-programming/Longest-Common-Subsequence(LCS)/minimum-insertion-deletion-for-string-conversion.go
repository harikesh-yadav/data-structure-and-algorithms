package main

import "fmt"

/**
problem:  https://www.geeksforgeeks.org/minimum-number-deletions-insertions-transform-one-string-another/

str1 = "heap", str2 = "pea"


time complexity:= O(n * m

space complexity:= O(n * m)

**/

func main() {
	str1 := "heap"
	str2 := "pea"
	n := len(str1)
	m := len(str2)

	delettion, insert := minimumInsertORDeleteForStringConversion(str1, str2, n, m)

	fmt.Println(delettion, insert)
}

func minimumInsertORDeleteForStringConversion(str1, str2 string, n, m int) (int, int) {
	dp := make([][]int, n+1)
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

	deletion := len(str1) - dp[n][m]
	insert := len(str2) - dp[n][m]
	return deletion, insert
}
