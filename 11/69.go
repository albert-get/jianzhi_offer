package main

import "fmt"

/**
山峰数组的顶部
*/
/**
题目：在一个长度大于或等于3的数组中，任意相邻的两个数字都不相等。该数组的前若干数字是递增的，之后的数字是递减的，因此它的值看起来像一座山峰。请找出山峰顶部，即数组中最大值的位置。例如，在数组[1，3，5，4，2]中，最大值是5，输出它在数组中的下标2。
*/

func peakIndexInMountainArray(nums []int) int {
	left := 1
	right := len(nums) - 2
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] > nums[mid-1] && nums[mid] > nums[mid+1] {
			return mid
		}
		if nums[mid] > nums[mid-1] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
func main() {
	p := peakIndexInMountainArray([]int{1, 3, 5, 4, 2})
	fmt.Println(p)
}
