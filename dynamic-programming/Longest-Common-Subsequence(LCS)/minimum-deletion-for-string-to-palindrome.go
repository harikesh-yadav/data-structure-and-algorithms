package main

import "fmt"

/**
Problem:  https://www.geeksforgeeks.org/minimum-number-deletions-make-string-palindrome/

Input : aebcbda
Output : 2
Remove characters 'e' and 'd'
Resultant string will be 'abcba'
which is a palindromic string
Input : geeksforgeeks
Output : 8


**/

func main() {
	str1 := "geeksforgeeks"
	n := len(str1)

	count := minimumDeletionForStringToPolidromic(str1, n)

	fmt.Println(count)
}

func minimumDeletionForStringToPolidromic(str1 string, n int) int {
	str2 := reverse((str1))
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

	return n - dp[n][n]

}

func reverse(str string) (s string) {
	for _, v := range str {
		s = string(v) + s
	}
	return
}
