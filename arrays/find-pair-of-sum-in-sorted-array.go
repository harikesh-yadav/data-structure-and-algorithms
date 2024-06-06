package main

import "fmt"

/**

https://www.geeksforgeeks.org/two-pointers-technique/


***/

func main() {
	arr := []int{5, 8, 10, 12}
	sum := 18
	ok := findPairOfSum(arr, sum)
	fmt.Println(ok)
}

func findPairOfSum(arr []int, sum int) bool {
	i := 0
	j := len(arr) - 1

	for i < j {
		if (arr[i] + arr[j]) == sum {
			return true
		}
		if (arr[i] + arr[j]) < sum {
			i++
		} else {
			j--
		}
	}
	return false
}
