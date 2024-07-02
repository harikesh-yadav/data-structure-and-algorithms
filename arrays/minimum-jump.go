package main

import (
	"fmt"
	"math"
)

func main() {
	nums := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 1, 0}

	result := jump(nums)
	fmt.Println(result)
}

func jump(nums []int) int {
	count := 0
	max := len(nums) - 1

	if len(nums) == 0 || nums[0] <= 0 {
		return 0
	}
	if len(nums) == 1 {
		return 0
	}

	for i := 0; i <= max; {
		fmt.Println("i:", i)
		if i+nums[i] >= max {
			count++
			break
		} else {
			count++
			i = maxIndex(nums, i, max)
		}
	}
	return count
}

func maxIndex(nums []int, i, max int) int {
	end := i + nums[i]
	start := i + 1
	min := math.MaxInt
	index := start
	for j := start; j <= end && j <= max; j++ {
		if min >= (max - (j + nums[j])) {
			min = max - (j + nums[j])
			index = j
		}
	}
	return index
}
