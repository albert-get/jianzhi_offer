package main

import (
	"fmt"
	"harry.com/jianzhi_offer/dui"
	"harry.com/jianzhi_offer/shu"
	"math"
)

/**
二叉树中每层的最大值
*/

/**
题目：输入一棵二叉树，请找出二叉树中每层的最大值。例如，输入图7.4中的二叉树，返回各层节点的最大值[3，4，9]。
*/

func largestValues(root *shu.TreeNode) (result []int) {
	current := 0
	next := 0
	queue := dui.Queue{}
	if root != nil {
		queue.Push(root)
		current = 1
	}
	max := math.MinInt
	for !queue.IsEmpty() {
		node, _ := queue.Pop().(*shu.TreeNode)
		current--
		max = int(math.Max(float64(max), float64(node.Val)))
		if node.Left != nil {
			queue.Push(node.Left)
			next++
		}
		if node.Right != nil {
			queue.Push(node.Right)
			next++
		}
		if current == 0 {
			result = append(result, max)
			max = math.MinInt
			current = next
			next = 0
		}
	}
	return result
}
func largestValues2(root *shu.TreeNode) (result []int) {
	queue1 := dui.Queue{}
	queue2 := dui.Queue{}
	if root != nil {
		queue1.Push(root)
	}
	max := math.MinInt
	for !queue1.IsEmpty() {
		node, _ := queue1.Pop().(*shu.TreeNode)
		max = int(math.Max(float64(max), float64(node.Val)))
		if node.Left != nil {
			queue2.Push(node.Left)
		}
		if node.Right != nil {
			queue2.Push(node.Right)
		}
		if queue1.IsEmpty() {
			result = append(result, max)
			max = math.MinInt
			queue1 = queue2
			queue2 = dui.Queue{}
		}
	}
	return result
}
func main() {
	tree := shu.ConstructorC2(
		[]shu.ToTree{
			{Val: 5, Level: 3},
			{Val: 4, Level: 2},
			{Val: 1, Level: 3},
			{Val: 3, Level: 1},
			{Val: 2, Level: 2},
			{Val: 9, Level: 3},
		}, 1)
	result := largestValues(tree)
	result2 := largestValues2(tree)
	fmt.Println(result)
	fmt.Println(result2)
}
