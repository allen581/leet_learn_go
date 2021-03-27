// code 0001

package array

import (
	"fmt"
	"leetcode/common"
)

// TwoSum 在数组中找到 2 个数之和等于给定值的数字，结果返回 2 个数字在数组中的下标
// 对每一个元素，计算其另一半数值，循环查询另一半是否在数组中，查到则返回下标
func TwoSum(nums []int, target int) []int {
	fmt.Println(nums)
	fmt.Println(target)
	results := []int{}
	for i := 0; i < len(nums); i++ {
		num1 := nums[i]
		num2 := target - num1
		index, found := common.FindInt(nums, num2)
		if found && i != index {
			results = append(results, i)
			results = append(results, index)
			return results
		}
	}
	return results
}

// TwoSum2 在数组中找到 2 个数之和等于给定值的数字，结果返回 2 个数字在数组中的下标
// 对每一个元素，计算其另一半数值，判断另一半是否在map中，如果在返回下标，不在则把该元素存入map中
func TwoSum2(nums []int, target int) []int {
	temp_map := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		num1 := nums[i]
		num2 := target - num1
		if _, ok := temp_map[num2]; ok {
			return []int{temp_map[num2], i}
		}
		temp_map[num1] = i
	}
	return nil
}
