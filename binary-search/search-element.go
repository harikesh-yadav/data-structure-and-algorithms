package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 8, 9}
	n := 8

	index := searchElement(nums, n)
	fmt.Println(index)
}

func searchElement(nums []int, n int) int {

	start := 0
	end := len(nums) - 1

	for start <= end {
		mid := start + (end-start)/2

		if nums[mid] == n {
			return mid
		}
		if mid < n {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1

}
