package main

import (
	"fmt"
	"harry.com/jianzhi_offer/set"
)

/**
包含重复元素集合的全排列
*/
/**
题目：给定一个包含重复数字的集合，请找出它的所有全排列。例如，集合[1，1，2]有3个全排列，分别是[1，1，2]、[1，2，1]和[2，1，1]。
*/

func permuteUnique(nums []int) [][]int {
	result := new([][]int)
	helper5(nums, 0, result)
	return *result
}
func helper5(nums []int, i int, result *[][]int) {
	if i == len(nums) {
		var permutation []int
		for _, num := range nums {
			permutation = append(permutation, num)
		}
		*result = append(*result, permutation)
	} else {
		myset := set.NewSetInt()
		for j := i; j < len(nums); j++ {
			if !myset.Contains(nums[j]) {
				myset.Add(nums[j])
				nums[i], nums[j] = nums[j], nums[i]
				helper5(nums, i+1, result)
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

}
func main() {
	r := permuteUnique([]int{1, 1, 2})
	fmt.Println(r)
}
