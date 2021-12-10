package main

import (
	"fmt"
	"math"
)

/**
最长递增路径
*/
/**
题目：输入一个整数矩阵，请求最长递增路径的长度。矩阵中的路径沿着上、下、左、右4个方向前行。例如，图15.14中矩阵的最长递增路径的长度为4，其中一条最长的递增路径为3→4→5→8，
*/

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	lengths := make([][]int, len(matrix))
	for i, _ := range lengths {
		lengths[i] = make([]int, len(matrix[0]))
	}
	longest := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			length := dfs3(matrix, &lengths, i, j)
			longest = int(math.Max(float64(longest), float64(length)))
		}
	}
	return longest
}
func dfs3(matrix [][]int, lengths *[][]int, i int, j int) int {
	if t := *lengths; t[i][j] != 0 {
		return t[i][j]
	}
	rows := len(matrix)
	cols := len(matrix[0])
	dirs := [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	length := 1
	for _, dir := range dirs {
		r := i + dir[0]
		c := j + dir[1]
		if r >= 0 && r < rows && c >= 0 && c < cols && matrix[r][c] > matrix[i][j] {
			path := dfs3(matrix, lengths, r, c)
			length = int(math.Max(float64(path+1), float64(length)))
		}
	}
	t := *lengths
	t[i][j] = length
	return length
}
func main() {
	r := longestIncreasingPath([][]int{{3, 4, 5}, {3, 2, 8}, {2, 2, 1}})
	fmt.Println(r)
}
