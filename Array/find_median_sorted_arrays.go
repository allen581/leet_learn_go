// code 0004

package array

// FindMedianSortedArrays1 合并数组后取中位值
// 时间复杂度为O(m+n)，不符合O(log (m+n))
func FindMedianSortedArrays1(nums1 []int, nums2 []int) float64 {
	length1 := len(nums1)
	length2 := len(nums2)
	total_length := length1 + length2
	var total_nums = make([]int, total_length)
	position1 := 0
	position2 := 0
	for position1 < length1 && position2 < length2 {
		temp_num1 := nums1[position1]
		temp_num2 := nums2[position2]
		if temp_num1 < temp_num2 {
			total_nums[position1+position2] = temp_num1
			position1 = position1 + 1
		} else {
			total_nums[position1+position2] = temp_num2
			position2 = position2 + 1
		}
	}
	for position1 < length1 {
		total_nums[position1+position2] = nums1[position1]
		position1 = position1 + 1
	}
	for position2 < length2 {
		total_nums[position1+position2] = nums2[position2]
		position2 = position2 + 1
	}
	result := float64(total_nums[(total_length-1)/2]+total_nums[total_length/2]) / 2.0
	return result
}

// findMedianSortedArrays2 不用合并数组
// 先算出中位数下标之和，然后维护分别指向num1和num2的指针，每次移动较小值，直到到达中位数的位置
func findMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	length1 := len(nums1)
	length2 := len(nums2)
	total_length := length1 + length2

}
