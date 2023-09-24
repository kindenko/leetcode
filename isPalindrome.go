package main

import (
	"strconv"
)

func IsPalindrome(x int) bool {
	stringValue := strconv.Itoa(x)
	n := len(stringValue)

	for i := 0; i < n/2; i++ {
		if stringValue[i] != stringValue[n-i-1] {
			return false
		}
	}

	return true

}
