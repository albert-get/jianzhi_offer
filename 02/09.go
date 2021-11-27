package main

import "fmt"

/**
乘积小于k的子数组
*/
/**
题目：输入一个由正整数组成的数组和一个正整数k
，请问数组中有多少个数字乘积小于k
的连续子数组？例如，输入数组[10，5，2，6]，k
的值为100，有8个子数组的所有数字的乘积小于100，它们分别是[10]、[5]、[2]、[6]、[10，5]、[5，2]、[2，6]和[5，2，6]。
*/

func numSubarrayProductLessThanK(nums []int, k int) int {
	left := 0
	right := 0
	temp := 1
	count := 0
	for ; right < len(nums); right++ {
		temp *= nums[right]
		for temp >= k && left <= right {
			temp /= nums[left]
			left++
		}
		if right >= left {
			count += right - left + 1
		}
	}
	return count
}
func main() {
	fmt.Print(numSubarrayProductLessThanK([]int{10, 5, 2, 6}, 100))
}
