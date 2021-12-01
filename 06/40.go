package main

import (
	"fmt"
	"harry.com/jianzhi_offer/zhan"
	"math"
)

/**
矩阵中的最大矩形
*/
/**
题目：请在一个由0、1组成的矩阵中找出最大的只包含1的矩形并输出它的面积。例如，在图6.6的矩阵中，最大的只包含1的矩阵如阴影部分所示，它的面积是6。
*/

func largestRectangleArea5(heights []int) int {
	stack := zhan.NewStackInt()
	stack.Push(-1)
	maxArea := 0
	for i := 0; i < len(heights); i++ {
		for stack.Peek() != -1 && heights[stack.Peek()] >= heights[i] {
			height := heights[stack.Pop()]
			width := i - stack.Peek() - 1
			maxArea = int(math.Max(float64(maxArea), float64(height*width)))
		}
		stack.Push(i)
	}
	for stack.Peek() != -1 {
		height := heights[stack.Pop()]
		width := len(heights) - stack.Peek() - 1
		maxArea = int(math.Max(float64(maxArea), float64(height*width)))
	}
	return maxArea
}
func maximalRectangle(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	heights := make([]int, len(matrix[0]))
	maxArea := 0
	for _, row := range matrix {
		for i := 0; i < len(row); i++ {
			if row[i] == 0 {
				heights[i] = 0
			} else {
				heights[i]++
			}
		}
		maxArea = int(math.Max(float64(maxArea), float64(largestRectangleArea5(heights))))
	}
	return maxArea
}
func main() {
	fmt.Println(maximalRectangle([][]int{{1, 0, 1, 0, 0}, {0, 0, 1, 1, 1}, {1, 1, 1, 1, 1}, {1, 0, 0, 1, 0}}))
}
