package main

import (
	"fmt"
	"harry.com/jianzhi_offer/shu"
	"harry.com/jianzhi_offer/zhan"
)

/**
展平二叉搜索树
*/
/**
题目：给定一棵二叉搜索树，请调整节点的指针使每个节点都没有左子节点。调整之后的树看起来像一个链表，但仍然是二叉搜索树。
*/

func increasingBST(root *shu.TreeNode) *shu.TreeNode {
	stack := zhan.NewStack()
	cur := root
	var prev *shu.TreeNode
	var first *shu.TreeNode
	for cur != nil || !stack.IsEmpty() {
		for cur != nil {
			stack.Push(cur)
			cur = cur.Left
		}
		node, _ := stack.Pop().(*shu.TreeNode)
		if prev != nil {
			prev.Right = node
		} else {
			first = node
		}
		prev = node
		node.Left = nil
		cur = node.Right
	}
	return first
}
func main() {
	tree := shu.ConstructorC2(
		[]shu.ToTree{
			{Val: 1, Level: 3},
			{Val: 2, Level: 2},
			{Val: 3, Level: 3},
			{Val: 4, Level: 1},
			{Val: 5, Level: 2},
			{Val: 6, Level: 3},
		}, 1)
	result := increasingBST(tree)
	fmt.Println(result)
}
