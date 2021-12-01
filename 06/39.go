package main

import (
	"fmt"
	"harry.com/jianzhi_offer/zhan"
	"math"
)

/**
直方图最大矩形面积
*/
/**
题目：直方图是由排列在同一基线上的相邻柱子组成的图形。输入一个由非负数组成的数组，数组中的数字是直方图中柱子的高。求直方图中最大矩形面积。假设直方图中柱子的宽都为1。
*/

func largestRectangleArea(heights []int) int {
	maxArea := 0
	for i := 0; i < len(heights); i++ {
		min := heights[i]
		for j := i; j < len(heights); j++ {
			min = int(math.Min(float64(min), float64(heights[j])))
			area := min * (j - i + 1)
			maxArea = int(math.Max(float64(maxArea), float64(area)))
		}
	}
	return maxArea
}
func largestRectangleArea2(heights []int) int {
	return helper(heights, 0, len(heights))
}
func helper(heights []int, start int, end int) int {
	if start == end {
		return 0
	}
	if start+1 == end {
		return heights[start]
	}
	minIndex := start
	for i := start + 1; i < end; i++ {
		if heights[i] < heights[minIndex] {
			minIndex = i
		}
	}
	area := (end - start) * heights[minIndex]
	left := helper(heights, start, minIndex)
	right := helper(heights, minIndex+1, end)
	area = int(math.Max(float64(area), float64(left)))
	return int(math.Max(float64(area), float64(right)))
}
func LargestRectangleArea3(heights []int) int {
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
func main() {
	fmt.Println(largestRectangleArea([]int{3, 2, 5, 4, 6, 1, 4, 2}))
	fmt.Println(largestRectangleArea2([]int{3, 2, 5, 4, 6, 1, 4, 2}))
	fmt.Println(LargestRectangleArea3([]int{3, 2, 5, 4, 6, 1, 4, 2}))
}
