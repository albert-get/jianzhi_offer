package main

import (
	"fmt"
	"harry.com/jianzhi_offer/dui"
	"harry.com/jianzhi_offer/shu"
)

/**
二叉树最低层最左边的值
*/

/**
题目：如何在一棵二叉树中找出它最低层最左边节点的值？假设二叉树中最少有一个节点。例如，在如图7.5所示的二叉树中最低层最左边一个节点的值是5。
*/

func findBottomLeftValue(root *shu.TreeNode) (bottomLeft int) {
	queue1 := dui.Queue{}
	queue2 := dui.Queue{}
	queue1.Push(root)
	bottomLeft = root.Val
	for !queue1.IsEmpty() {
		node, _ := queue1.Pop().(*shu.TreeNode)
		if node.Left != nil {
			queue2.Push(node.Left)
		}
		if node.Right != nil {
			queue2.Push(node.Right)
		}
		if queue1.IsEmpty() {
			queue1 = queue2
			queue2 = dui.Queue{}
			if !queue1.IsEmpty() {
				bottomLeftNode, _ := queue1.Peek().(*shu.TreeNode)
				bottomLeft = bottomLeftNode.Val
			}
		}
	}
	return bottomLeft
}
func main() {
	tree := shu.ConstructorC2(
		[]shu.ToTree{
			{Val: 5, Level: 3},
			{Val: 6, Level: 2},
			{Val: 7, Level: 3},
			{Val: 8, Level: 1},
			{Val: 9, Level: 3},
			{Val: 10, Level: 2},
			{Val: 11, Level: 3},
		}, 1)
	result := findBottomLeftValue(tree)
	fmt.Println(result)
}
