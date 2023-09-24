package main

func TwoSum(nums []int, target int) []int {

	b := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if val, ok := b[nums[i]]; ok {
			return []int{val, i}
		}
		b[target-nums[i]] = i

	}
	return []int{}
}
