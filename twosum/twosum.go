package main

import "fmt"

func twoSum(nums []int, target int) []int {
	cnt := len(nums)
	fmt.Println(cnt)
	var tmp []int

	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			fmt.Println(i, j, nums[i], nums[i])
			if target == nums[i]+nums[j] {
				tmp[0] = i
				tmp[1] = j
				return tmp
			}
		}
	}
	return nums
}

func main() {
	nums := []int{2, 7, 9, 14}
	target := 9

	result := twoSum(nums, target)
	fmt.Println(result[0], result[1])
}
