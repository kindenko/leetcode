package main

func PlusOne(digits []int) []int {

	for i := len(digits) - 1; i >= 0; i-- {

		if digits[i]+1 < 10 {
			digits[i] = digits[i] + 1
			return digits
		} else {
			digits[i] = 0
		}

	}
	if digits[0] == 0 {
		return append([]int{1}, digits...)
	}

	return digits
}
