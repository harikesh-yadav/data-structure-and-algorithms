package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

/*****
Problems:   https://www.geeksforgeeks.org/find-next-greater-number-set-digits/


Input:  n = "218765"
Output: "251678"
*****/

func main() {
	input := "218765"

	result := solve(input)

	fmt.Print(result)
}

func solve(input string) []int {

	strArr := strings.Split(input, "")
	digitsArr := convert(strArr)

	if len(digitsArr) <= 1 {
		return []int{-1}
	}

	var i int
	for i = len(digitsArr) - 2; i > 0; i-- {
		if digitsArr[i] < digitsArr[i+1] {
			break
		}
	}

	if i == 0 {
		return []int{-1}
	}

	min := math.MaxInt
	var k int
	for j := i + 1; j <= len(strArr)-1; j++ {
		if digitsArr[j] < min {
			k = j
			min = digitsArr[j]
		}
	}

	digitsArr[i], digitsArr[k] = digitsArr[k], digitsArr[i]

	remains := digitsArr[i+1:]
	sort.Ints(remains)

	return digitsArr

}

func convert(str []string) []int {
	arrInt := make([]int, len(str))

	for i, s := range str {
		v, _ := strconv.Atoi(s)
		arrInt[i] = v
	}
	return arrInt
}
