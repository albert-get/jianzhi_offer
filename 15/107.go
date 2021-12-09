package main

import (
	"fmt"
	"harry.com/jianzhi_offer/dui"
	"math"
)

/**
矩阵中的距离
*/
/**
题目：输入一个由0、1组成的矩阵M
，请输出一个大小相同的矩阵D
，矩阵D
中的每个格子是矩阵M
中对应格子离最近的0的距离。水平或竖直方向相邻的两个格子的距离为1。假设矩阵M
中至少有一个0。
*/

func updateMatrix(matrix [][]int) [][]int {
	rows := len(matrix)
	cols := len(matrix[0])
	dists := make([][]int, rows)
	for i, _ := range dists {
		dists[i] = make([]int, cols)
	}
	queue := dui.Queue{}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == 0 {
				queue.Push([]int{i, j})
				dists[i][j] = 0
			} else {
				dists[i][j] = math.MaxInt
			}
		}
	}
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for queue.Len() > 0 {
		pos := queue.Pop().([]int)
		dist := dists[pos[0]][pos[1]]
		for _, dir := range dirs {
			r := pos[0] + dir[0]
			c := pos[1] + dir[1]
			if r >= 0 && c >= 0 && r < rows && c < cols {
				if dists[r][c] > dist+1 {
					dists[r][c] = dist + 1
					queue.Push([]int{r, c})
				}
			}
		}
	}
	return dists
}
func main() {
	r := updateMatrix([][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}})
	fmt.Println(r)
}
