package main

import "fmt"

func main() {

	str1 := "AXY"
	str2 := "ADXCPY"
	n := len(str1)
	m := len(str2)

	ok := isSubSequence(str1, str2, n, m)
	fmt.Println(ok)

}

func isSubSequence(str1, str2 string, n, m int) bool {

	var i, j int
	for i < n && j < m {
		if n == m {
			if str1[i] == str2[j] {
				i++
				j++
			} else {
				return false
			}
		} else if n < m {
			if str1[i] == str2[j] {
				i++
				j++
			} else {
				j++
			}
		} else {
			if str1[i] == str2[j] {
				i++
				j++
			} else {
				i++
			}
		}
	}

	if n < m {
		return i == n
	}
	return j == m

}
