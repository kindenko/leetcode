package main

import "fmt"

func removeElement(nums []int, val int) int {

	k := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			continue
		}
		fmt.Println(nums[i])
		nums[k] = nums[i]
		k++
	}

	return k
}
