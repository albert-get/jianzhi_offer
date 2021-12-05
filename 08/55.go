package main

import (
	"fmt"
	"harry.com/jianzhi_offer/shu"
	"harry.com/jianzhi_offer/zhan"
)

/**
二叉搜索树迭代器
*/
/**
题目：请实现二叉搜索树的迭代器BSTIterator，它主要有如下3个函数。
·构造函数：输入二叉搜索树的根节点初始化该迭代器。
·函数next：返回二叉搜索树中下一个最小的节点的值。
·函数hasNext：返回二叉搜索树是否还有下一个节点。
*/

type BSTIterator struct {
	cur   *shu.TreeNode
	stack *zhan.Stack
}

func Constructor(root *shu.TreeNode) BSTIterator {
	return BSTIterator{
		cur:   root,
		stack: zhan.NewStack(),
	}
}
func (bs *BSTIterator) hasNext() bool {
	if bs.cur != nil || !bs.stack.IsEmpty() {
		return true
	} else {
		return false
	}
}
func (bs *BSTIterator) next() int {
	for bs.cur != nil {
		bs.stack.Push(bs.cur)
		bs.cur = bs.cur.Left
	}
	node, _ := bs.stack.Pop().(*shu.TreeNode)
	val := node.Val
	bs.cur = node.Right
	return val
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
	bs := Constructor(tree)
	fmt.Println(bs.hasNext())
	fmt.Println(bs.next())
}
