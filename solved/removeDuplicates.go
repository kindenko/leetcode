package main

import "fmt"

func removeDuplicates(nums []int) int {
	l := len(nums)

	if l == 1 {
		return 1
	}

	k := 0

	for i := 1; i < l; i++ {
		if nums[k] != nums[i] {
			k++

			nums[k] = nums[i]
		}
	}

	fmt.Println(nums)
	return k + 1
}
