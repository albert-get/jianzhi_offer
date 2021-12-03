package main

import (
	"fmt"
	"harry.com/jianzhi_offer/dui"
	"harry.com/jianzhi_offer/shu"
)

/**
二叉树的右侧视图
*/
/**
题目：给定一棵二叉树，如果站在该二叉树的右侧，那么从上到下看到的节点构成二叉树的右侧视图。
*/

func rightSideView(root *shu.TreeNode) (view []int) {
	if root == nil {
		return view
	}
	queue1 := dui.Queue{}
	queue2 := dui.Queue{}
	queue1.Push(root)
	for !queue1.IsEmpty() {
		node, _ := queue1.Pop().(*shu.TreeNode)
		if node.Left != nil {
			queue2.Push(node.Left)
		}
		if node.Right != nil {
			queue2.Push(node.Right)
		}
		if queue1.IsEmpty() {
			view = append(view, node.Val)
			queue1 = queue2
			queue2 = dui.Queue{}
		}
	}
	return view
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
	result := rightSideView(tree)
	fmt.Println(result)
}
