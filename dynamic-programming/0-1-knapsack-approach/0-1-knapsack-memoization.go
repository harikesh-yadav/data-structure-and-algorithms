package main

import "fmt"

/*
Input: N = 3, W = 4, profit[] = {1, 2, 3}, weight[] = {4, 5, 1}
Output: 3


Time Complexity: O(N * W). As redundant calculations of states are avoided.
Auxiliary Space: O(N * W) + O(N). The use of a 2D array data structure for storing intermediate states and O(N) auxiliary stack space(ASS) has been used for recursion stack

*/

var DP [][]int

func main() {
	value := []int{1, 2, 3}
	weight := []int{4, 5, 1}

	size := 3
	W := 3
	InitDP(size, W) // Init matrix
	profit := KnapsackMemoization(value, weight, size, W)
	fmt.Println(profit)
}

func KnapsackMemoization(value, weight []int, size, W int) int {

	if size == 0 || W == 0 {
		return DP[size][W]
	}

	if DP[size][W] != -1 {
		return DP[size][W]
	}
	if weight[size-1] < W {
		p1 := value[size-1] + KnapsackMemoization(value, weight, size-1, W-weight[size-1])
		p2 := KnapsackMemoization(value, weight, size-1, W)
		if p1 > p2 {
			DP[size][W] = p1
		} else {
			DP[size][W] = p2
		}
	} else {
		DP[size][W] = KnapsackMemoization(value, weight, size-1, W)
	}
	return DP[size][W]
}

func InitDP(size, W int) {
	DP = make([][]int, size+1)
	for i := 0; i < size+1; i++ {
		DP[i] = make([]int, W+1)
	}
	for i := 0; i < size+1; i++ {
		for j := 0; j < W+1; j++ {
			if i == 0 || j == 0 {
				DP[i][j] = 0
			} else {
				DP[i][j] = -1
			}
		}
	}
}
