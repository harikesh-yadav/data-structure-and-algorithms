package main

import "fmt"

/*
Problem:  https://www.geeksforgeeks.org/0-1-knapsack-problem-dp-10/
Input: N = 3, W = 4, profit[] = {1, 2, 3}, weight[] = {4, 5, 1}
Output: 3


Time Complexity: O(N * W). where ‘N’ is the number of elements and ‘W’ is capacity.
Auxiliary Space: O(N * W). The use of a 2-D array of size ‘N*W’.
*/

func main() {
	value := []int{20, 30, 10, 50}
	weight := []int{1, 3, 4, 6}

	size := len(value)
	W := 10
	profit := KnapsackDPMatrix(value, weight, size, W)

	fmt.Println(profit)
}

func KnapsackDPMatrix(value, weight []int, size, W int) int {
	dp := make([][]int, size+1)
	for i := 0; i < size+1; i++ {
		dp[i] = make([]int, W+1)
	}

	for i := 1; i < size+1; i++ {
		for w := 1; w < W+1; w++ {
			if weight[i-1] <= w {
				p1 := value[i-1] + dp[i-1][w-weight[i-1]]
				p2 := dp[i-1][w]
				if p1 > p2 {
					dp[i][w] = p1
				} else {
					dp[i][w] = p2
				}
			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
	}

	// for i := 0; i < size+1; i++ {
	// 	for j := 0; j < W+1; j++ {
	// 		fmt.Print(dp[i][j], ", ")
	// 	}
	// 	fmt.Println()
	// }

	//  Print set for weight consider

	i := size
	j := W
	set := make([]int, 0)
	result := dp[size][W]
	for i > 0 && j > 0 {
		if result != dp[i-1][j] {
			result = result - value[i-1]
			set = append(set, weight[i-1])
			j = j - weight[i-1]
		}
		i--
	}
	fmt.Println(set)

	return dp[size][W]
}
