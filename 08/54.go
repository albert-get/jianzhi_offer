package main

import (
	"fmt"
	"harry.com/jianzhi_offer/shu"
	"harry.com/jianzhi_offer/zhan"
)

/**
所有大于或等于节点的值之和
*/
/**
题目：给定一棵二叉搜索树，请将它的每个节点的值替换成树中大于或等于该节点值的所有节点值之和。假设二叉搜索树中节点的值唯一。
*/

func convertBST(root *shu.TreeNode) *shu.TreeNode {
	stack := zhan.NewStack()
	cur := root
	sum := 0
	for cur != nil || !stack.IsEmpty() {
		for cur != nil {
			stack.Push(cur)
			cur = cur.Right
		}
		node, _ := stack.Pop().(*shu.TreeNode)
		sum += node.Val
		node.Val = sum
		cur = node.Left
	}
	return root
}
func main() {
	tree := shu.ConstructorC2(
		[]shu.ToTree{
			{Val: 1, Level: 3},
			{Val: 2, Level: 2},
			{Val: 3, Level: 3},
			{Val: 4, Level: 1},
			{Val: 5, Level: 3},
			{Val: 6, Level: 2},
			{Val: 7, Level: 3},
		}, 1)
	result := convertBST(tree)
	fmt.Println(result)
}
