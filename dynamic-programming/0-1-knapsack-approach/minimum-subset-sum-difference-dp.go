package main

import (
	"fmt"
)

/*
Input:  arr[] = {1, 6, 11, 5}
Output: 1
Explanation:
Subset1 = {1, 5, 6}, sum of Subset1 = 12
Subset2 = {11}, sum of Subset2 = 11

time complexity O(n*sum)
space complexity O(n*sum)
*/

var matrix [][]bool

func main() {
	arr := []int{1, 6, 11, 5}
	size := len(arr)
	sum := arraySum(arr)
	diff := minimumSubsetSumDifferenceDP(arr, size, sum)
	fmt.Println(diff)
}

func minimumSubsetSumDifferenceDP(arr []int, size, sum int) int {
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
	min := 0
	for i := sum / 2; i >= 0; i-- {
		if matrix[size][i] {
			min = sum - 2*i
			break
		}
	}
	return min
}

func arraySum(arr []int) int {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	return sum
}
