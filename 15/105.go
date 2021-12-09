package main

import (
	"fmt"
	"harry.com/jianzhi_offer/dui"
	"harry.com/jianzhi_offer/zhan"
	"math"
)

/**
最大的岛屿
*/
/**
题目：海洋岛屿地图可以用由0、1组成的二维数组表示，水平或竖直方向相连的一组1表示一个岛屿，请计算最大的岛屿的面积（即岛屿中1的数目）。例如，在图15.5中有4个岛屿，其中最大的岛屿的面积为5。
*/

func maxAreaOfIsland(grid [][]int) int {
	rows := len(grid)
	cols := len(grid[0])
	visited := make([][]bool, rows)
	for i, _ := range visited {
		visited[i] = make([]bool, cols)
	}
	maxArea := 0
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if grid[i][j] == 1 && !visited[i][j] {
				area := getArea(grid, visited, i, j)
				//area := getArea1(grid, visited, i, j)
				//area := getArea2(grid, visited, i, j)
				maxArea = int(math.Max(float64(area), float64(maxArea)))
			}
		}
	}
	return maxArea
}
func getArea(grid [][]int, visited [][]bool, i int, j int) int {
	queue := dui.Queue{}
	queue.Push([]int{i, j})
	visited[i][j] = true
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	area := 0
	for queue.Len() != 0 {
		pos, _ := queue.Pop().([]int)
		area++
		for _, dir := range dirs {
			r := pos[0] + dir[0]
			c := pos[1] + dir[1]
			if r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0]) && grid[r][c] == 1 && !visited[r][c] {
				queue.Push([]int{r, c})
				visited[r][c] = true
			}
		}
	}
	return area
}
func getArea1(grid [][]int, visited [][]bool, i int, j int) int {
	stack := zhan.NewStack()
	stack.Push([]int{i, j})
	visited[i][j] = true
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	area := 0
	for !stack.IsEmpty() {
		pos, _ := stack.Pop().([]int)
		area++

		for _, dir := range dirs {
			r := pos[0] + dir[0]
			c := pos[1] + dir[1]
			if r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0]) && grid[r][c] == 1 && !visited[r][c] {
				stack.Push([]int{r, c})
				visited[r][c] = true
			}
		}
	}
	return area
}
func getArea2(grid [][]int, visited [][]bool, i int, j int) int {
	area := 1
	visited[i][j] = true
	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, dir := range dirs {
		r := i + dir[0]
		c := j + dir[1]
		if r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0]) && grid[r][c] == 1 && !visited[r][c] {
			area += getArea2(grid, visited, r, c)
		}
	}
	return area
}
func main() {
	r := maxAreaOfIsland([][]int{{1, 1, 0, 0, 1}, {1, 0, 0, 1, 0}, {1, 1, 0, 1, 0}, {0, 0, 1, 0, 0}})
	fmt.Println(r)
}
