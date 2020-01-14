package main

import "fmt"

func twoSum(nums []int, target int) []int {
	cnt := len(nums)

	for i := 0; i < cnt; i++ {
		for j := i + 1; j < cnt; j++ {
			if target == nums[i]+nums[j] {
				nums[0] = i
				nums[1] = j
				return nums[:2]
			}
		}
	}
	return nums
}

func main() {
	nums := []int{2, 0, 9, 14}
	target := 9

	result := twoSum(nums, target)
	fmt.Println(result[0], result[1])
}
