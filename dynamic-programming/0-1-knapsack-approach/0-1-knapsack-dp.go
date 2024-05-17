package main

import "fmt"

/*
Input: N = 3, W = 4, profit[] = {1, 2, 3}, weight[] = {4, 5, 1}
Output: 3


Time Complexity: O(N * W). where ‘N’ is the number of elements and ‘W’ is capacity.
Auxiliary Space: O(N * W). The use of a 2-D array of size ‘N*W’.
*/

var matrix [][]int

func main() {
	value := []int{1, 2, 3}
	weight := []int{4, 5, 1}

	size := 3
	W := 3
	InitMatrix(size, W) // init matrix Matrix
	profit := KnapsackDPMatrix(value, weight, size, W)

	fmt.Println(profit)
}

func KnapsackDPMatrix(value, weight []int, size, W int) int {
	if size == 0 || W == 0 {
		return matrix[size][W]
	}

	for i := 1; i < size+1; i++ {
		for w := 1; w < W+1; w++ {
			if weight[i-1] <= w {
				p1 := value[i-1] + matrix[i-1][w-weight[i-1]]
				p2 := matrix[i-1][w]
				if p1 > p2 {
					matrix[i][w] = p1
				} else {
					matrix[i][w] = p2
				}
			} else {
				matrix[i][w] = matrix[i-1][w]
			}
		}
	}
	return matrix[size][W]
}

func InitMatrix(size, W int) {
	matrix = make([][]int, size+1)
	for i := 0; i < size+1; i++ {
		matrix[i] = make([]int, W+1)
	}
	for i := 0; i < size+1; i++ {
		for j := 0; j < W+1; j++ {
			if i == 0 || j == 0 {
				matrix[i][j] = 0
			} else {
				matrix[i][j] = -1
			}
		}
	}
}
