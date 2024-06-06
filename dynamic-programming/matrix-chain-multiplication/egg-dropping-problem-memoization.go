package main

import (
	"fmt"
	"math"
)

/*
*****
Problem:  https://www.geeksforgeeks.org/egg-dropping-puzzle-dp-11/

Time Complexity: As there is a case of overlapping sub-problems the time complexity is exponential.
Auxiliary Space: O(1). As there was no use of any data structure for storing values.
*****
*/
var matrix [][]int

func main() {
	egg := 2
	floor := 10

	initMatrix(egg, floor)

	result := eggDropProblemRecursiveMemozation(egg, floor)
	fmt.Println(result)
}

func eggDropProblemRecursiveMemozation(egg, floor int) int {
	if floor == 0 || floor == 1 {
		return floor
	}
	if egg == 1 {
		return floor
	}

	if matrix[egg][floor] != -1 {
		return matrix[egg][floor]
	}

	min := math.MaxInt

	for k := 1; k <= floor; k++ {
		// use maximum for worst case
		temp := 1 + maximum(eggDropProblemRecursiveMemozation(egg-1, k-1), eggDropProblemRecursiveMemozation(egg, floor-k))

		if min > temp {
			min = temp
		}
	}
	matrix[egg][floor] = min

	return matrix[egg][floor]

}

func maximum(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func initMatrix(egg, floor int) {
	matrix = make([][]int, egg+1)

	for m := 0; m < egg+1; m++ {
		matrix[m] = make([]int, floor+1)
	}

	for m := 0; m < egg+1; m++ {
		for n := 0; n < floor+1; n++ {
			matrix[m][n] = -1
		}
	}
}
