package main

import (
	"fmt"
	"math"
)

/*
******

https://www.geeksforgeeks.org/palindrome-partitioning-dp-17/

*******
*/
var matrix [][]int

func main() {

	str := "ababbbabbababa"
	i := 0
	j := len(str) - 1
	initMatrix(j)
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
	if matrix[i][j] != -1 {
		return matrix[i][j]
	}
	min := math.MaxInt
	for k := i; k < j; k++ {
		var left, right int
		if matrix[i][k] != -1 {
			left = matrix[i][k]
		} else {
			left = polidromePartitionRecursive(str, i, k)
			matrix[i][k] = left
		}
		if matrix[k+1][j] != -1 {
			right = matrix[k+1][j]
		} else {
			right = polidromePartitionRecursive(str, k+1, j)
			matrix[k+1][j] = right
		}
		temp := left + right + 1
		if min > temp {
			min = temp
		}
	}
	matrix[i][j] = min
	return matrix[i][j]

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

func initMatrix(size int) {
	matrix = make([][]int, size+1)

	for m := 0; m < size+1; m++ {
		matrix[m] = make([]int, size+1)
	}

	for m := 0; m < size+1; m++ {
		for n := 0; n < size+1; n++ {
			matrix[m][n] = -1
		}

	}
}
