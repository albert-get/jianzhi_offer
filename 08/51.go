package main

import (
	"fmt"
	"harry.com/jianzhi_offer/shu"
	"math"
)

/**
节点值之和最大的路径
*/
/**
题目：在二叉树中将路径定义为顺着节点之间的连接从任意一个节点开始到达任意一个节点所经过的所有节点。路径中至少包含一个节点，不一定经过二叉树的根节点，也不一定经过叶节点。给定非空的一棵二叉树，请求出二叉树所有路径上节点值之和的最大值。
*/

func maxPathSum(root *shu.TreeNode) int {
	maxSum := []int{math.MinInt}
	dfs3(root, maxSum)
	return maxSum[0]
}
func dfs3(root *shu.TreeNode, maxSum []int) int {
	if root == nil {
		return 0
	}
	maxSumLeft := []int{math.MinInt}
	left := int(math.Max(0, float64(dfs3(root.Left, maxSumLeft))))
	maxSumRight := []int{math.MinInt}
	right := int(math.Max(0, float64(dfs3(root.Right, maxSumRight))))

	maxSum[0] = int(math.Max(float64(maxSumLeft[0]), float64(maxSumRight[0])))
	maxSum[0] = int(math.Max(float64(maxSum[0]), float64(root.Val+left+right)))
	return root.Val + int(math.Max(float64(left), float64(right)))
}
func main() {
	tree := shu.ConstructorC2(
		[]shu.ToTree{
			{Val: 4, Level: 2},
			{Val: -9, Level: 1},
			{Val: -3, Level: 4},
			{Val: 15, Level: 3},
			{Val: 20, Level: 2},
			{Val: 7, Level: 3},
		}, 1)
	sum:=maxPathSum(tree)
	fmt.Println(sum)
}
