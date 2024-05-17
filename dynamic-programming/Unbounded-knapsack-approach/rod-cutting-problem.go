package main

import "fmt"

/*
https://www.geeksforgeeks.org/cutting-a-rod-dp-13/

length   | 1   2   3   4   5   6   7   8
--------------------------------------------
price    | 1   5   8   9  10  17  17  20

time complexity:  O(n^2)
space complexity: O(n^2)
*/

func main() {
	length := []int{1, 2, 3, 4, 5, 6, 7, 8}
	price := []int{1, 5, 8, 9, 10, 17, 17, 20}
	size := len(price)
	sum := 8

	profit := rodCutting(price, length, size, sum)
	fmt.Println(profit)
}

func rodCutting(value, lenght []int, size, sum int) int {
	dp := make([][]int, size+1)
	for i := 0; i < size+1; i++ {
		dp[i] = make([]int, sum+1)
	}

	for i := 1; i < size+1; i++ {
		for j := 1; j < sum+1; j++ {
			if lenght[i-1] <= j {
				p1 := value[i-1] + dp[i][j-lenght[i-1]]
				p2 := dp[i-1][j]
				if p1 > p2 {
					dp[i][j] = p1
				} else {
					dp[i][j] = p2
				}
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}

	return dp[size][sum]

}
