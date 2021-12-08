package main

import (
	"fmt"
	"math"
)

/**
最小路径之和
*/
/**
题目：在一个m×n（m、n均大于0）的格子中，每个位置都有一个数字。一个机器人每步只能向下或向右，请计算它从格子的左上角到达右下角的路径的数字之和的最小值。例如，从图14.8中3×3的格子的左上角到达右下角的路径的数字之和的最小值是8，图中数字之和最小的路径用灰色背景表示。
*/

func minPathSum(grid [][]int) int {
	dp := make([][]int, len(grid))
	for i, _ := range dp {
		dp[i] = make([]int, len(grid[0]))
	}
	dp[0][0] = grid[0][0]
	for j := 1; j < len(grid[0]); j++ {
		dp[0][j] = grid[0][j] + dp[0][j-1]
	}
	for i := 1; i < len(grid); i++ {
		dp[i][0] = grid[i][0] + dp[i-1][0]
		for j := 1; j < len(grid[0]); j++ {
			prev := int(math.Min(float64(dp[i-1][j]), float64(dp[i][j-1])))
			dp[i][j] = grid[i][j] + prev
		}
	}
	return dp[len(grid)-1][len(grid[0])-1]
}

func minPathSum1(grid [][]int) int {
	dp := make([]int, len(grid[0]))
	dp[0] = grid[0][0]
	for j := 1; j < len(grid[0]); j++ {
		dp[j] = grid[0][j-1] + dp[j-1]
	}
	for i := 1; i < len(grid); i++ {
		dp[0] += grid[i][0]
		for j := 1; j < len(grid[0]); j++ {
			dp[j] = grid[i][j] + int(math.Min(float64(dp[j]), float64(dp[j-1])))
		}
	}
	return dp[len(grid[0])-1]
}
func main() {
	//r:=minPathSum([][]int{{1,3,1},{2,5,2},{3,4,1}})
	r := minPathSum1([][]int{{1, 3, 1}, {2, 5, 2}, {3, 4, 1}})
	fmt.Println(r)
}
