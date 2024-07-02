package main

import "fmt"

/***
Problem:  https://www.geeksforgeeks.org/find-a-peak-in-a-given-array/

**/

func main() {
	nums := []int{10, 20, 15, 2, 23, 90, 67}

	index := findPeekElement(nums)
	fmt.Println(index)

}

func findPeekElement(nums []int) int {
	start := 0
	size := len(nums)
	end := size - 1

	for start <= end {
		mid := start + (end-start)/2

		prev := (mid - 1 + size) % size
		next := (mid + 1) % size

		if nums[prev] <= nums[mid] && nums[mid] >= nums[next] {
			return mid
		} else if nums[mid] <= nums[next] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}
