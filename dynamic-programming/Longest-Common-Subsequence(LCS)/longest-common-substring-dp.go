package main

import "fmt"

/*
problem:  https://www.geeksforgeeks.org/longest-common-substring-dp-29/
Input : X = “abcdxyz”, y = “xyzabcd”
Output : 4
Explanation:
The longest common substring is “abcd” and is of length 4.

time complexity:  O(n * m)
space complexity:  O(n * m)

*/

func main() {
	str1 := "abcdxyz"
	str2 := "xyzabcd"
	n := len(str1)
	m := len(str2)

	count := longestCommonSubstring(str1, str2, n, m)

	fmt.Println(count)
}

func longestCommonSubstring(str1, str2 string, n, m int) int {
	max := 0
	dp := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		dp[i] = make([]int, m+1)
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if str1[i-1] == str2[j-1] {
				dp[i][j] = 1 + dp[i-1][j-1]
				if max < dp[i][j] {
					max = dp[i][j]
				}
			} else {
				dp[i][j] = 0
			}
		}
	}
	return max
}
