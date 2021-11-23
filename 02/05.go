package main

import "fmt"

/**
和为k的子数组
*/

/**
题目：输入一个整数数组和一个整数k
，请问数组中有多少个数字之和等于k
的连续子数组？例如，输入数组[1，1，1]，k
的值为2，有2个连续子数组之和等于2。
*/

func subArraySum(nums []int, k int) int {
	sumToCount := make(map[int]int) //go里面初始化一般都用make
	sumToCount[0] = 1
	sum := 0
	count := 0
	for _, num := range nums {
		sum += num
		count += sumToCount[sum-k]
		sumToCount[sum] += 1
	}
	return count
}
func main() {
	fmt.Print(subArraySum([]int{1, 1, 1}, 2))
}
