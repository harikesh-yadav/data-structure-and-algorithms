package main

import (
	"fmt"
)

/*
Problem:  https://www.geeksforgeeks.org/longest-common-subsequence-dp-4/

Input: S1 = “AGGTAB”, S2 = “GXTXAYB”
Output: 4

time complexity:  O(2^(n+m))
space complexity: O(1)
*/
func main() {
	str1 := "BD"
	str2 := "ABCD"
	n := len(str1)
	m := len(str2)

	count := longestCommonSubSequenceRecursion(str1, str2, n, m)

	fmt.Println(count)

}

func longestCommonSubSequenceRecursion(str1, str2 string, n, m int) int {
	if n == 0 || m == 0 {
		return 0
	}
	if str1[n-1] == str2[m-1] {
		return longestCommonSubSequenceRecursion(str1, str2, n-1, m-1) + 1
	}

	s1 := longestCommonSubSequenceRecursion(str1, str2, n-1, m)
	s2 := longestCommonSubSequenceRecursion(str1, str2, n, m-1)

	if s1 > s2 {
		return s1
	}
	return s2
}
