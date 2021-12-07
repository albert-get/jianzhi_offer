package main

import "fmt"

/**
允许重复选择元素的组合
*/
/**
题目：给定一个没有重复数字的正整数集合，请找出所有元素之和等于某个给定值的所有组合。同一个数字可以在组合中出现任意次。例如，输入整数集合[2，3，5]，元素之和等于8的组合有3个，分别是[2，2，2，2]、[2，3，3]和[3，5]。
*/

func combinationSum(nums []int, target int) [][]int {
	result := new([][]int)
	//combination := new([]int)
	var combination []int
	helper2(nums, target, 0, combination, result)
	return *result
}
func helper2(nums []int, target int, i int, combination []int, result *[][]int) {
	if target == 0 {
		*result = append(*result, combination)
	} else if target > 0 && i < len(nums) {
		helper2(nums, target, i+1, combination, result)

		combination = append(combination, nums[i])

		helper2(nums, target-nums[i], i, combination, result)

		combination = combination[:len(combination)-1]
	}
}
func main() {
	r := combinationSum([]int{2,3,5}, 8)
	fmt.Println(r)
}
