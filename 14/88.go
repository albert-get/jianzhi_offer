package main

import (
	"fmt"
	"math"
)

/**
爬楼梯的最少成本
*/
/**
题目：一个数组cost的所有数字都是正数，它的第i
个数字表示在一个楼梯的第i
级台阶往上爬的成本，在支付了成本cost[i
]之后可以从第i
级台阶往上爬1级或2级。假设台阶至少有2级，既可以从第0级台阶出发，也可以从第1级台阶出发，请计算爬上该楼梯的最少成本。例如，输入数组[1，100，1，1，100，1]，则爬上该楼梯的最少成本是4，分别经过下标为0、2、3、5的这4级台阶，如图14.1所示。
*/

func minCostClimbingStairs(cost []int) int {
	leng := len(cost)
	return int(math.Min(float64(helper10(cost, leng-2)), float64(helper10(cost, leng-1))))
}
func helper10(cost []int, i int) int {
	if i < 2 {
		return cost[i]
	}
	return int(math.Min(float64(helper10(cost, i-2)), float64(helper10(cost, i-1)))) + cost[i]
}
func minCostClimbingStairs1(cost []int) int {
	l := len(cost)
	dp := make([]int, l)
	helper11(cost, l-1, dp)
	return int(math.Min(float64(dp[l-2]), float64(dp[l-1])))
}
func helper11(cost []int, i int, dp []int) {
	if i < 2 {
		dp[i] = cost[i]
	} else if dp[i] == 0 {
		helper11(cost, i-2, dp)
		helper11(cost, i-1, dp)
		dp[i] = int(math.Min(float64(dp[i-2]), float64(dp[i-1]))) + cost[i]
	}
}

func minCostClimbingStairs2(cost []int) int {
	l := len(cost)
	dp := make([]int, l)
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < l; i++ {
		dp[i] = int(math.Min(float64(dp[i-2]), float64(dp[i-1]))) + cost[i]
	}
	return int(math.Min(float64(dp[l-2]), float64(dp[l-1])))
}
func minCostClimbingStairs3(cost []int) int {
	dp := []int{cost[0], cost[1]}
	for i := 2; i < len(cost); i++ {
		dp[i%2] = int(math.Min(float64(dp[0]), float64(dp[1]))) + cost[i]
	}
	return int(math.Min(float64(dp[0]), float64(dp[1])))
}
func main() {
	//c := minCostClimbingStairs([]int{1, 100, 1, 1, 100, 1})
	//c := minCostClimbingStairs1([]int{1, 100, 1, 1, 100, 1})
	//c := minCostClimbingStairs2([]int{1, 100, 1, 1, 100, 1})
	c := minCostClimbingStairs3([]int{1, 100, 1, 1, 100, 1})
	fmt.Println(c)
}
