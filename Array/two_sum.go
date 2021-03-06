package array

import (
	"fmt"
	"leetcode/common"
)

// TwoSum 在数组中找到 2 个数之和等于给定值的数字，结果返回 2 个数字在数组中的下标
func TwoSum(nums []int, target int) []int {
	fmt.Println(nums)
	fmt.Println(target)
	results := []int{}
	for i := 0; i < len(nums); i++ {
		num1 := nums[i]
		num2 := target - num1
		index, found := common.FindInt(nums, num2)
		if found && i != index{
			results = append(results, i)
			results = append(results, index)
			return results
		}
	}
	return results
}

// result := Array.TwoSum([]int{2, 7, 11, 15}, 9)
// fmt.Println(result)
