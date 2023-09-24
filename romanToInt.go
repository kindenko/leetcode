package main

import (
	"fmt"
)

func RomanToInt(s string) int {
	symboles := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	num := 0

	if len(s) == 1 {
		num += symboles[string(s[0])]
		return num
	}

	for i := 0; i < len(s); i++ {

		if symboles[string(s[i])] < symboles[string(s[i+1])] {
			num += symboles[string(s[i+1])] - symboles[string(s[i])]
			i += 1
		} else {
			num += symboles[string(s[i])]
		}

		if i == len(s)-2 {
			num += symboles[string(s[len(s)-1])]
			fmt.Println(symboles[string(s[len(s)-1])])
			return num
		}

	}

	return num
}
