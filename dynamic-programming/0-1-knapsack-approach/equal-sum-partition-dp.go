package main

import (
	"fmt"
)

/*
Input: N = 4
arr = {1, 5, 11, 5}
Output: YES
Explanation:
The two parts are {1, 5, 5} and {11}.
*/
var matrix [][]bool

func main() {
	arr := []int{1, 5, 11, 5}
	size := 4

	ok := isEqualSumSubset(arr, size)
	fmt.Println(ok)
}

func isEqualSumSubset(arr []int, size int) bool {
	sum := arraySum(arr)
	if sum%2 != 0 {
		return false
	}
	sum = sum / 2

	matrix = make([][]bool, size+1)
	for i := 0; i < size+1; i++ {
		matrix[i] = make([]bool, sum+1)
	}

	for i := 0; i < size+1; i++ {
		for j := 0; j < sum+1; j++ {
			if i == 0 {
				matrix[i][j] = false
			}
			if j == 0 {
				matrix[i][j] = true
			}
		}
	}

	for i := 1; i < size+1; i++ {
		for j := 1; j < sum+1; j++ {
			if arr[i-1] <= j {
				matrix[i][j] = matrix[i-1][j-arr[i-1]] || matrix[i-1][j]
			} else {
				matrix[i][j] = matrix[i-1][j]
			}
		}
	}
	return matrix[size][sum]

}

func arraySum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}
