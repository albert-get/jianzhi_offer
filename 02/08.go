package main

import (
	"fmt"
	"math"
)

/**
和大于或等于k的最短子数组
*/

/**
题目：输入一个正整数组成的数组和一个正整数k
，请问数组中和大于或等于k
的连续子数组的最短长度是多少？如果不存在所有数字之和大于或等于k
的子数组，则返回0。例如，输入数组[5，1，4，3]，k
的值为7，和大于或等于7的最短连续子数组是[4，3]，因此输出它的长度2。
*/

func minSubArrayLen(k int, nums []int) int {
	left := 0
	right := 0
	sum := 0
	minLength := math.MaxInt
	for ; right < len(nums); right++ {
		sum += nums[right]
		for left <= right && sum >= k {
			minLength = int(math.Min(float64(minLength), float64(right-left+1)))
			sum -= nums[left]
			left++
		}
	}
	if minLength == math.MaxInt {
		return 0
	} else {
		return minLength
	}
}
func main() {
	fmt.Print(minSubArrayLen(7, []int{5, 1, 4, 3}))
}
