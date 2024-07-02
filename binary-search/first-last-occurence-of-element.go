package main

import "fmt"

func main() {
	nums := []int{2, 4, 10, 10, 10, 15, 18}
	n := 10

	result1 := firstOccurence(nums, n)
	fmt.Println(result1)

	result2 := lastOccurence(nums, n)

	fmt.Println(result2)
}

func firstOccurence(nums []int, n int) int {
	start := 0
	end := len(nums) - 1
	ans := -1

	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] == n {
			ans = mid
			end = mid - 1

		} else if nums[mid] > n {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return ans
}

func lastOccurence(nums []int, n int) int {
	start := 0
	end := len(nums) - 1
	ans := -1

	for start <= end {
		mid := start + (end-start)/2

		if nums[mid] == n {
			ans = mid
			start = mid + 1
		} else if nums[mid] > n {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return ans
}
