package main

import (
	"fmt"
	"harry.com/jianzhi_offer/shu"
	"harry.com/jianzhi_offer/zhan"
)

/**
二叉搜索树的下一个节点
*/
/**
题目：给定一棵二叉搜索树和它的一个节点p
，请找出按中序遍历的顺序该节点p
的下一个节点。假设二叉搜索树中节点的值都是唯一的。
*/

func inorderSuccessor(root *shu.TreeNode, p *shu.TreeNode) *shu.TreeNode {
	stack := zhan.NewStack()
	cur := root
	found := false
	for cur != nil || !stack.IsEmpty() {
		for cur != nil {
			stack.Push(cur)
			cur = cur.Left
		}
		node, _ := stack.Pop().(*shu.TreeNode)
		if found {
			cur = node
			break
		} else if p == node {
			found = true
		}
		cur = node.Right

	}
	return cur
}
func inorderSuccessor2(root *shu.TreeNode, p *shu.TreeNode) *shu.TreeNode {
	cur := root
	var result *shu.TreeNode
	for cur != nil {
		if cur.Val > p.Val {
			result = cur
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	return result
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
	p := tree
	//next := inorderSuccessor(tree, p)
	next := inorderSuccessor2(tree, p)
	fmt.Println(next)
}
