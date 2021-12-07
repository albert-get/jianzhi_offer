package main

import (
	"fmt"
	"math"
)

/**
环形房屋偷盗
*/

/**
题目：一条环形街道上有若干房屋。输入一个数组表示该条街道上的房屋内财产的数量。如果这条街道上相邻的两幢房屋被盗就会自动触发报警系统。请计算小偷在这条街道上最多能偷取的财产的数量。例如，街道上5家的财产用数组[2，3，4，5，3]表示，如果小偷到下标为1和3的房屋内盗窃，那么他能偷取到价值为8的财物，这是他在不触发报警系统的情况下能偷取到的最多的财物，如图14.4所示。被盗的房屋上方用特殊符号标出。
*/

func rob5(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	result1 := help13(nums, 0, len(nums)-2)
	result2 := help13(nums, 1, len(nums)-1)
	return int(math.Max(float64(result1), float64(result2)))
}
func help13(nums []int, start int, end int) int {
	dp := [2]int{}
	dp[0] = nums[start]
	if start < end {
		dp[1] = int(math.Max(float64(nums[start]), float64(nums[start+1])))
	}
	for i := start + 2; i <= end; i++ {
		j := i - start
		dp[j%2] = int(math.Max(float64(dp[(j-1)%2]), float64(dp[(j-2)%2]+nums[i])))
	}
	return dp[(end-start)%2]
}
func main() {
	m := rob5([]int{2, 3, 4, 5, 3})
	fmt.Println(m)
}
