package main

import (
	"fmt"
	"math"
)

/**
粉刷房子
*/
/**
题目：一排n
幢房子要粉刷成红色、绿色和蓝色，不同房子被粉刷成不同颜色的成本不同。用一个n
×3的数组表示n
幢房子分别用3种颜色粉刷的成本。要求任意相邻的两幢房子的颜色都不一样，请计算粉刷这n
幢房子的最少成本。例如，粉刷3幢房子的成本分别为[[17，2，16]，[15，14，5]，[13，3，1]]，如果分别将这3幢房子粉刷成绿色、蓝色和绿色，那么粉刷的成本是10，是最少的成本。
*/

func minCost(costs [][]int) int {
	if len(costs) == 0 {
		return 0
	}
	dp := [3][2]int{}
	for j := 0; j < 3; j++ {
		dp[j][0] = costs[0][j]
	}
	for i := 1; i < len(costs); i++ {
		for j := 0; j < 3; j++ {
			prev1 := dp[(j+2)%3][(i-1)%2]
			prev2 := dp[(j+1)%3][(i-1)%2]
			dp[j][i%2] = int(math.Min(float64(prev1), float64(prev2))) + costs[i][j]
		}
	}
	last := (len(costs) - 1) % 2
	return int(math.Min(float64(dp[0][last]), float64(math.Min(float64(dp[1][last]), float64(dp[2][last])))))
}
func main() {
	c := minCost([][]int{{17, 2, 16}, {15, 14, 5}, {13, 3, 1}})
	fmt.Println(c)
}
