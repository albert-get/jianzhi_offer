package main

import "fmt"

/**
排序数组中只出现一次的数字
*/
/**
题目：在一个排序的数组中，除一个数字只出现一次之外，其他数字都出现了两次，请找出这个唯一只出现一次的数字。例如，在数组[1，1，2，2，3，4，4，5，5]中，数字3只出现了一次。
*/

func singleNonDuplicate(nums []int) int {
	left := 0
	right := len(nums) / 2
	for left <= right {
		mid := (left + right) / 2
		i := mid * 2
		if i < len(nums)-1 && nums[i] != nums[i+1] {
			if mid == 0 || nums[i-2] != nums[i-1] {
				return nums[i]
			}
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return nums[len(nums)-1]
}
func main() {
	s := singleNonDuplicate([]int{1, 1, 2, 2, 3, 4, 4, 5, 5})
	fmt.Println(s)
}
