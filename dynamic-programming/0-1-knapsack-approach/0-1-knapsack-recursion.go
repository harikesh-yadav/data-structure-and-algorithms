package main

import "fmt"

/*
Input: N = 3, W = 4, profit[] = {1, 2, 3}, weight[] = {4, 5, 1}
Output: 3


Time Complexity: O(2 pow N)
Auxiliary Space: O(N), Stack space required for recursion
*/

func main() {
	value := []int{1, 2, 3}
	weight := []int{4, 5, 1}
	size := 3
	W := 4

	profit := KnapsackRecursion(value, weight, size, W)
	fmt.Println(profit)
}

func KnapsackRecursion(value, weight []int, size, W int) int {
	if size == 0 || W == 0 {
		return 0
	}

	if weight[size-1] < W {
		p1 := value[size-1] + KnapsackRecursion(value, weight, size-1, W-weight[size-1])
		p2 := KnapsackRecursion(value, weight, size-1, W)
		if p1 > p2 {
			return p1
		} else {
			return p2
		}
	}
	return KnapsackRecursion(value, weight, size-1, W)
}
