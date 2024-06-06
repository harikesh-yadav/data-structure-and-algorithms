package main

import (
	"fmt"
	"math"
)

/******
Problem:  https://www.geeksforgeeks.org/egg-dropping-puzzle-dp-11/


Time Complexity: O(n*k*k) where n is number of eggs and k is number of floors.
Auxiliary Space: O(n*k) as 2D memoisation table has been used.
******/

func main() {
	egg := 2
	floor := 10
	result := eggDropProblemRecursive(egg, floor)
	fmt.Println(result)
}

func eggDropProblemRecursive(egg, floor int) int {
	if floor == 0 || floor == 1 {
		return floor
	}
	if egg == 1 {
		return floor
	}

	min := math.MaxInt

	for k := 1; k <= floor; k++ {
		// use maximum for worst case
		temp := 1 + maximum(eggDropProblemRecursive(egg-1, k-1), eggDropProblemRecursive(egg, floor-k))
		if min > temp {
			min = temp
		}
	}
	return min

}

func maximum(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}
