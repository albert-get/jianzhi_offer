package main

import (
	"fmt"
	"harry.com/jianzhi_offer/shu"
)

/**
从根节点到叶节点的路径数字之和
*/

/**
题目：在一棵二叉树中所有节点都在0～9的范围之内，从根节点到叶节点的路径表示一个数字。求二叉树中所有路径表示的数字之和。
*/

func sumNumbers(root *shu.TreeNode) int {
	return dfs1(root, 0)
}
func dfs1(root *shu.TreeNode, path int) int {
	if root == nil {
		return 0
	}
	path = path*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return path
	}
	return dfs1(root.Left, path) + dfs1(root.Right, path)
}
func main() {
	tree := shu.ConstructorC2(
		[]shu.ToTree{
			{Val: 5, Level: 3},
			{Val: 9, Level: 2},
			{Val: 1, Level: 3},
			{Val: 3, Level: 1},
			{Val: 0, Level: 2},
			{Val: 2, Level: 3},
		}, 1)
	sum := sumNumbers(tree)
	fmt.Println(sum)
}
