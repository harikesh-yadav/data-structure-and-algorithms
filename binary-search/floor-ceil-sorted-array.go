package main

import "fmt"

func main() {

	nums := []int{1, 2, 8, 10, 10, 12, 19}

	target := 5

	floor := getFloor(nums, target)
	ceil := getCeil(nums, target)

	fmt.Println(floor, ceil)

}

func getFloor(nums []int, target int) int {

	start := 0
	end := len(nums) - 1
	result := 0
	for start <= end {

		mid := start + (end-start)/2

		if target == nums[mid] {
			return nums[mid]
		}

		if nums[mid] < target {
			result = nums[mid]
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return result
}

func getCeil(nums []int, target int) int {

	start := 0
	end := len(nums) - 1
	result := 0
	for start <= end {
		mid := start + (end-start)/2

		if nums[mid] == target {
			return nums[mid]
		}

		if nums[mid] > target {
			result = nums[mid]
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return result
}
