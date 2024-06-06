package main

import (
	"fmt"
	"math"
)

/*
**
Problem:  https://www.geeksforgeeks.org/matrix-chain-multiplication-dp-8/

Input: arr[] = {40, 20, 30, 10, 30}
Output: 26000

Steps:-
Choose i and j position wisely
Choose invalid base condition based on problem
Find K look scheme
Calculate answer for temp answer
and return answer

time complexity : O(n)
space complexity : O(1)

**
*/
func main() {
	arr := []int{40, 20, 30, 10, 30}
	i := 1
	j := len(arr) - 1

	result := calculateMinMultiplicationCost(arr, i, j)
	fmt.Println(result)
}

func calculateMinMultiplicationCost(arr []int, i, j int) int {
	if i == j {
		return 0
	}
	min := math.MaxInt

	for k := i; k < j; k++ {
		temp := calculateMinMultiplicationCost(arr, i, k) + calculateMinMultiplicationCost(arr, k+1, j) + (arr[i-1] * arr[k] * arr[j])
		if temp < min {
			min = temp
		}
	}
	return min
}
