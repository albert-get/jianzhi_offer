package main

import (
	"fmt"
	"math"
)

/**
三角形中最小路径之和
*/
/**
题目：在一个由数字组成的三角形中，第1行有1个数字，第2行有2个数字，以此类推，第n
行有n
个数字。例如，图14.9是一个包含4行数字的三角形。如果每步只能前往下一行中相邻的数字，请计算从三角形顶部到底部的路径经过的数字之和的最小值。如图14.9所示，从三角形顶部到底部的路径数字之和的最小值为11，对应的路径经过的数字用阴影表示。
*/

func minimumTotal(triangle [][]int) int {
	size := len(triangle)
	dp := make([][]int, size)
	for i, _ := range dp {
		dp[i] = make([]int, size)
	}
	for i := 0; i < size; i++ {
		for j := 0; j <= i; j++ {
			dp[i][j] = triangle[i][j]
			if i > 0 && j == 0 {
				dp[i][j] += dp[i-1][j]
			} else if i > 0 && i == j {
				dp[i][j] += dp[i-1][j-1]
			} else if i > 0 {
				dp[i][j] += int(math.Min(float64(dp[i-1][j]), float64(dp[i-1][j-1])))
			}
		}
	}
	min := math.MaxInt
	for _, num := range dp[size-1] {
		min = int(math.Min(float64(min), float64(num)))
	}
	return min
}

func minimumTotal1(triangle [][]int) int {
	dp := make([]int, len(triangle))
	for _, row := range triangle {
		for j := len(row) - 1; j >= 0; j-- {
			if j == 0 {
				dp[j] += row[j]
			} else if j == len(row)-1 {
				dp[j] = dp[j-1] + row[j]
			} else {
				dp[j] = int(math.Min(float64(dp[j]), float64(dp[j-1]))) + row[j]
			}
		}
	}
	min := math.MaxInt
	for _, num := range dp {
		min = int(math.Min(float64(min), float64(num)))
	}
	return min
}
func main() {
	//r := minimumTotal([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}})
	r := minimumTotal1([][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}})
	fmt.Println(r)
}
