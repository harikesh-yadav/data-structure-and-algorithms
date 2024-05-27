package main

import "fmt"

/*
Input: set[] = {3, 34, 4, 12, 5, 2}, sum = 9
Output: True


time complexity O(size*sum)
space complexity   O(size*sum)

*/

var matrix [][]bool

func main() {
	arr := []int{3, 34, 4, 12, 5, 2}
	size := len(arr)
	sum := 9

	InitMMATRIX(size, sum)
	isExist := SubsetSumDP(arr, size, sum)
	fmt.Println(isExist)
}

func SubsetSumDP(arr []int, size, sum int) bool {
	if size == 0 || sum == 0 {
		return matrix[size][sum]
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

	// find subset also....
	subset := make([]int, 0)
	i := size
	j := sum
	for i > 0 && j > 0 {
		if !matrix[i-1][j] {
			subset = append(subset, arr[i-1])
			j -= arr[i-1]
		}
		i--

	}

	fmt.Println(subset, "----")
	return matrix[size][sum]
}

func InitMMATRIX(size, sum int) {
	matrix = make([][]bool, size+1)
	for i := 0; i < size+1; i++ {
		matrix[i] = make([]bool, sum+1)
	}

	for n := 0; n < size+1; n++ {
		for s := 0; s < sum+1; s++ {
			if n == 0 {
				matrix[n][s] = false
			}
			if s == 0 {
				matrix[n][s] = true
			}
		}
	}

}
