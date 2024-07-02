package main

import "fmt"

func main() {
	nums := []int{7, 9, 11, 12, 5}

	count := findRotationCount(nums)
	fmt.Println(count)
}

func findRotationCount(nums []int) int {
	start := 0
	size := len(nums)
	end := size - 1

	for start <= end {
		mid := start + (end-start)/2
		prev := (mid - 1 + size) % size
		next := (mid + 1) % size

		if nums[prev] >= nums[mid] && nums[mid] <= nums[next] {
			return mid
		}

		if nums[start] > nums[mid] {
			end = mid - 1
		}
		if nums[mid] > nums[end] {
			start = mid + 1
		}
	}
	return -1
}
