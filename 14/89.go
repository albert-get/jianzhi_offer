package main

import (
	"fmt"
	"math"
)

/**
房屋偷盗
*/
/**
题目：输入一个数组表示某条街道上的一排房屋内财产的数量。如果这条街道上相邻的两幢房屋被盗就会自动触发报警系统。请计算小偷在这条街道上最多能偷取到多少财产。例如，街道上5幢房屋内的财产用数组[2，3，4，5，3]表示，如果小偷到下标为0、2和4的房屋内盗窃，那么他能偷取到价值为9的财物，这是他在不触发报警系统的情况下能偷取到的最多的财物，如图14.3所示。被盗的房屋上方用特殊符号标出。
*/

func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	for i, _ := range dp {
		dp[i] = -1
	}
	helper12(nums, len(nums)-1, dp)
	return dp[len(nums)-1]
}
func helper12(nums []int, i int, dp []int) {
	if i == 0 {
		dp[i] = nums[0]
	} else if i == 1 {
		dp[i] = int(math.Max(float64(nums[0]), float64(nums[1])))
	} else if dp[i] < 0 {
		helper12(nums, i-2, dp)
		helper12(nums, i-1, dp)
		dp[i] = int(math.Max(float64(dp[i-1]), float64(dp[i-2]+nums[i])))
	}
}
func rob1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	if len(nums) > 1 {
		dp[1] = int(math.Max(float64(nums[0]), float64(nums[1])))
	}
	for i := 2; i < len(nums); i++ {
		dp[i] = int(math.Max(float64(dp[i-1]), float64(dp[i-2]+nums[i])))
	}
	return dp[len(nums)-1]
}
func rob2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, 2)
	dp[0] = nums[0]
	if len(nums) > 1 {
		dp[1] = int(math.Max(float64(nums[0]), float64(nums[1])))
	}
	for i := 2; i < len(nums); i++ {
		dp[i%2] = int(math.Max(float64(dp[(i-1)%2]), float64(dp[(i-2)%2]+nums[i])))
	}
	return dp[(len(nums)-1)%2]
}
func rob3(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	dp := [2][2]int{}
	dp[0][0] = 0
	dp[1][0] = nums[0]
	for i := 1; i < l; i++ {
		dp[0][i%2] = int(math.Max(float64(dp[0][(i-1)%2]), float64(dp[1][(i-1)%2])))
		dp[1][i%2] = nums[i] + dp[0][(i-1)%2]
	}
	return int(math.Max(float64(dp[0][(l-1)%2]), float64(dp[1][(l-1)%2])))
}
func main() {
	//m := rob([]int{2, 3, 4, 5, 3})
	//m := rob1([]int{2, 3, 4, 5, 3})
	//m := rob2([]int{2, 3, 4, 5, 3})
	m := rob3([]int{2, 3, 4, 5, 3})
	fmt.Println(m)
}
