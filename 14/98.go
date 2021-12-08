package main

import "fmt"

/**
路径的数目
*/
/**
题目：一个机器人从m
×n
的格子的左上角出发，它每步要么向下要么向右，直到抵达格子的右下角。请计算机器人从左上角到达右下角的路径的数目。例如，如果格子的大小是3×3，那么机器人从左上角到达右下角有6条符合条件的不同路径
*/

func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i, _ := range dp {
		dp[i] = make([]int, n)
	}
	return helper15(m-1, n-1, dp)
}
func helper15(i int, j int, dp [][]int) int {
	if dp[i][j] == 0 {
		if i == 0 || j == 0 {
			dp[i][j] = 1
		} else {
			dp[i][j] = helper15(i-1, j, dp) + helper15(i, j-1, dp)
		}
	}
	return dp[i][j]
}
func uniquePaths1(m int, n int) int {
	dp := make([][]int, m)
	for i, _ := range dp {
		dp[i] = make([]int, n)
	}
	for i, _ := range dp[0] {
		dp[0][i] = 1
	}
	for i := 1; i < m; i++ {
		dp[i][0] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}
	return dp[m-1][n-1]
}
func uniquePaths2(m int, n int) int {
	dp := make([]int, n)
	for i, _ := range dp {
		dp[i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] += dp[j-1]
		}
	}
	return dp[n-1]
}
func main() {
	//r:=uniquePaths(3,3)
	//r := uniquePaths1(3, 3)
	r := uniquePaths2(3, 3)
	fmt.Println(r)
}
