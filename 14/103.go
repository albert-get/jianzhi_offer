package main

import (
	"fmt"
	"math"
)

/**
最少的硬币数目
*/
/**
题目：给定正整数数组coins表示硬币的面额和一个目标总额t
，请计算凑出总额t
至少需要的硬币数目。每种硬币可以使用任意多枚。如果不能用输入的硬币凑出给定的总额，则返回-1。例如，如果硬币的面额为[1，3，9，10]，总额t
为15，那么至少需要3枚硬币，即2枚面额为3的硬币及1枚面额为9的硬币。
*/

func coinChange(coins []int, target int) int {
	dp := make([]int, target+1)
	for i, _ := range dp {
		dp[i] = target + 1
	}
	dp[0] = 0
	for _, coin := range coins {
		for j := target; j >= 1; j-- {
			for k := 1; k*coin <= j; k++ {
				dp[j] = int(math.Min(float64(dp[j]), float64(dp[j-k*coin]+k)))
			}
		}
	}
	if dp[target] > target {
		return -1
	} else {
		return dp[target]
	}
}
func coinChange1(coins []int, target int) int {
	dp := make([]int, target+1)
	for i := 1; i <= target; i++ {
		dp[i] = target + 1
		for _, coin := range coins {
			if i >= coin {
				dp[i] = int(math.Min(float64(dp[i]), float64(dp[i-coin]+1)))
			}
		}
	}
	if dp[target] > target {
		return -1
	} else {
		return dp[target]
	}
}
func main() {
	//r := coinChange([]int{1, 3, 9, 10}, 15)
	r := coinChange1([]int{1, 3, 9, 10}, 15)
	fmt.Println(r)
}
