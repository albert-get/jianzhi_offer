package main

import "fmt"

/**
：查找插入位置
*/
/**
题目：输入一个排序的整数数组nums和一个目标值t
，如果数组nums中包含t
，则返回t
在数组中的下标；如果数组nums中不包含t
，则返回将t
按顺序插入数组nums中的下标。假设数组中没有相同的数字。例如，输入数组nums为[1，3，6，8]，如果目标值t
为3，则输出1；如果t
为5，则返回2。
*/

func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] >= target {
			if mid == 0 || nums[mid-1] < target {
				return mid
			}
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return len(nums)
}
func main() {
	t := searchInsert([]int{1, 3, 6, 8}, 2)
	fmt.Println(t)
}
