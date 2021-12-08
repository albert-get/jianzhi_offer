package main

import "fmt"

/**
排列的数目
*/
/**
题目：给定一个非空的正整数数组nums和一个目标值t
，数组中的所有数字都是唯一的，请计算数字之和等于t
的所有排列的数目。数组中的数字可以在排列中出现任意次。例如，输入数组[1，2，3]，目标值t
为3，那么总共有4个组合的数字之和等于3，它们分别为{1，1，1}、{1，2}、{2，1}及{3}。
*/

func permutationSum(nums []int, target int) int {
	dp := make([]int, target+1)
	dp[0] = 1
	for i := 1; i <= target; i++ {
		for _, num := range nums {
			if i >= num {
				dp[i] += dp[i-num]
			}
		}
	}
	return dp[target]
}
func main() {
	r := permutationSum([]int{1, 2, 3}, 3)
	fmt.Println(r)
}
