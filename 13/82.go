package main

import (
	"fmt"
	"sort"
)

/**
包含重复元素集合的组合
*/
/**
题目：给定一个可能包含重复数字的整数集合，请找出所有元素之和等于某个给定值的所有组合。输出中不得包含重复的组合。例如，输入整数集合[2，2，2，4，3，3]，元素之和等于8的组合有2个，分别是[2，2，4]和[2，3，3]。
*/

func combinationSum2(nums []int, target int) [][]int {
	sort.Ints(nums)
	result := new([][]int)
	var combination []int
	helper4(nums, target, 0, combination, result)
	return *result
}
func helper4(nums []int, target int, i int, combination []int, result *[][]int) {
	if target == 0 {
		*result = append(*result, combination)
	} else if target > 0 && i < len(nums) {
		helper4(nums, target, getNext(nums, i), combination, result)
		combination = append(combination, nums[i])
		helper4(nums, target-nums[i], i+1, combination, result)
		combination = combination[:len(combination)-1]
	}
}
func getNext(nums []int, index int) int {
	next := index
	for next < len(nums) && nums[next] == nums[index] {
		next++
	}
	return next
}

func main() {
	r := combinationSum2([]int{2, 2, 2, 4, 3, 3}, 8)
	fmt.Println(r)
}
