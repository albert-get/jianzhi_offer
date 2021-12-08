package main

import "fmt"

/**
分割等和子集
*/
/**
题目：给定一个非空的正整数数组，请判断能否将这些数字分成和相等的两部分。例如，如果输入数组为[3，4，1]，将这些数字分成[3，1]和[4]两部分，它们的和相等，因此输出true；如果输入数组为[1，2，3，5]，则不能将这些数字分成和相等的两部分，因此输出false。
*/

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 == 1 {
		return false
	}
	//return subsetSum(nums, sum/2)
	//return subsetSum1(nums, sum/2)
	return subsetSum2(nums, sum/2)
}
func subsetSum(nums []int, target int) bool {
	dp := make([][]*bool, len(nums)+1)
	for i, _ := range dp {
		dp[i] = make([]*bool, target+1)
	}
	return *helper16(nums, dp, len(nums), target)
}
func helper16(nums []int, dp [][]*bool, i int, j int) *bool {
	if dp[i][j] == nil {
		if j == 0 {
			t := true
			dp[i][j] = &t
		} else if i == 0 {
			t := false
			dp[i][j] = &t
		} else {
			dp[i][j] = helper16(nums, dp, i-1, j)
			if !*dp[i][j] && j >= nums[i-1] {
				dp[i][j] = helper16(nums, dp, i-1, j-nums[i-1])
			}
		}
	}
	return dp[i][j]
}

func subsetSum1(nums []int, target int) bool {
	dp := make([][]bool, len(nums)+1)
	for i, _ := range dp {
		dp[i] = make([]bool, target+1)
	}
	for i, _ := range nums {
		dp[i][0] = true
	}
	for i := 1; i <= len(nums); i++ {
		for j := 1; j <= target; j++ {
			dp[i][j] = dp[i-1][j]
			if !dp[i][j] && j >= nums[i-1] {
				dp[i][j] = dp[i-1][j-nums[i-1]]
			}
		}
	}
	return dp[len(nums)][target]
}
func subsetSum2(nums []int, target int) bool {
	dp := make([]bool, target+1)
	dp[0] = true
	for i := 1; i <= len(nums); i++ {
		for j := target; j > 0; j-- {
			if !dp[j] && j >= nums[i-1] {
				dp[j] = dp[j-nums[i-1]]
			}
		}
	}
	return dp[target]
}
func main() {
	r := canPartition([]int{3, 4, 1})
	fmt.Println(r)
}
