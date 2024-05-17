package main

import (
	"fmt"
)

/*
Problem:  https://www.geeksforgeeks.org/longest-common-subsequence-dp-4/

Input: S1 = “AGGTAB”, S2 = “GXTXAYB”
Output: 4

time complexity:  O(n * m))
space complexity: O(n * m)
*/

var matrix [][]int

func main() {
	str1 := "BD"
	str2 := "ABCD"
	n := len(str1)
	m := len(str2)

	initMatrix(n, m) // Initilize matrix with -1 for each cell

	count := longestCommonSubSequenceMemoization(str1, str2, n, m)

	fmt.Println(count)

}

func longestCommonSubSequenceMemoization(str1, str2 string, n, m int) int {
	if n == 0 || m == 0 {
		return 0
	}

	if matrix[n][m] != -1 {
		return matrix[n][m]
	}

	if str1[n-1] == str2[m-1] {
		matrix[n][m] = longestCommonSubSequenceMemoization(str1, str2, n-1, m-1) + 1
	} else {
		s1 := longestCommonSubSequenceMemoization(str1, str2, n-1, m)
		s2 := longestCommonSubSequenceMemoization(str1, str2, n, m-1)
		if s1 > s2 {
			matrix[n][m] = s1
		} else {
			matrix[n][m] = s2
		}
	}
	return matrix[n][m]
}

func initMatrix(n, m int) {
	matrix = make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		matrix[i] = make([]int, m+1)
	}
	for i := 0; i < n+1; i++ {
		for j := 0; j < m+1; j++ {
			matrix[i][j] = -1
		}
	}
}
