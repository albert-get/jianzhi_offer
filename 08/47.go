package main

import (
	"fmt"
	"harry.com/jianzhi_offer/shu"
)

/**
二叉树剪枝
*/
/**
题目：一棵二叉树的所有节点的值要么是0要么是1，请剪除该二叉树中所有节点的值全都是0的子树。
*/

func pruneTree(root *shu.TreeNode) *shu.TreeNode {
	if root == nil {
		return root
	}
	root.Left = pruneTree(root.Left)
	root.Right = pruneTree(root.Right)
	if root.Left == nil && root.Right == nil && root.Val == 0 {
		return nil
	}
	return root
}
func main() {
	tree := shu.ConstructorC2(
		[]shu.ToTree{
			{Val: 5, Level: 3},
			{Val: 6, Level: 2},
			{Val: 7, Level: 3},
			{Val: 8, Level: 1},
			{Val: 10, Level: 2},
		}, 1)
	fmt.Println(shu.InorderTraversalD(tree))
	fmt.Println(shu.InorderTraversalT(tree))
	fmt.Println(shu.PreorderTraversalD(tree))
	fmt.Println(shu.PreorderTraversalT(tree))
	fmt.Println(shu.PostorderTraversalD(tree))
	fmt.Println(shu.PostorderTraversalT(tree))
	nodeTree := shu.ConstructorC2(
		[]shu.ToTree{
			{Val: 0, Level: 3},
			{Val: 0, Level: 2},
			{Val: 0, Level: 3},
			{Val: 1, Level: 1},
			{Val: 0, Level: 3},
			{Val: 0, Level: 2},
			{Val: 1, Level: 3},
		}, 1)
	fmt.Println(nodeTree)
	nodeTree = pruneTree(nodeTree)
	fmt.Println(nodeTree)
}
