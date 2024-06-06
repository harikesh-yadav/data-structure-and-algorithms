package main

import (
	"fmt"
	"math"
)

/*******

https://www.geeksforgeeks.org/palindrome-partitioning-dp-17/

Time Complexity: O(2n)
Auxiliary Space: O(n)

********/

func main() {

	str := "ababbbabbababa"
	i := 0
	j := len(str) - 1

	result := polidromePartitionRecursive(str, i, j)

	fmt.Println(result)
}

func polidromePartitionRecursive(str string, i, j int) int {
	if i >= j {
		return 0
	}
	if isPolidrome(str, i, j) {
		return 0
	}
	min := math.MaxInt
	for k := i; k < j; k++ {
		temp := polidromePartitionRecursive(str, i, k) + polidromePartitionRecursive(str, k+1, j) + 1
		if min > temp {
			min = temp
		}
	}
	return min

}

func isPolidrome(str string, i, j int) bool {
	for i < j {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}
	return true
}
