package main

import (
	"fmt"
	"harry.com/jianzhi_offer/set"
	"harry.com/jianzhi_offer/shu"
	"harry.com/jianzhi_offer/zhan"
)

/**
二叉搜索树中两个节点的值之和
*/
/**
题目：给定一棵二叉搜索树和一个值k
，请判断该二叉搜索树中是否存在值之和等于k
的两个节点。假设二叉搜索树中节点的值均唯一。
*/

func findTarget(root *shu.TreeNode, k int) bool {
	set1 := set.NewSetInt()
	stack := zhan.NewStack()
	cur := root
	for cur != nil || !stack.IsEmpty() {
		for cur != nil {
			stack.Push(cur)
			cur = cur.Left
		}
		node, _ := stack.Pop().(*shu.TreeNode)
		if set1.Contains(k - node.Val) {
			return true
		}
		set1.Add(node.Val)
		cur = node.Right
	}
	return false
}

type BSTIterator1 struct {
	cur   *shu.TreeNode
	stack *zhan.Stack
}

func Constructor1(root *shu.TreeNode) BSTIterator1 {
	return BSTIterator1{
		cur:   root,
		stack: zhan.NewStack(),
	}
}
func (bs *BSTIterator1) hasNext() bool {
	if bs.cur != nil || !bs.stack.IsEmpty() {
		return true
	} else {
		return false
	}
}
func (bs *BSTIterator1) next() int {
	for bs.cur != nil {
		bs.stack.Push(bs.cur)
		bs.cur = bs.cur.Left
	}
	node, _ := bs.stack.Pop().(*shu.TreeNode)
	val := node.Val
	bs.cur = node.Right
	return val
}

type BSTIteratorReversed struct {
	cur   *shu.TreeNode
	stack *zhan.Stack
}

func constructorBR(root *shu.TreeNode) (br BSTIteratorReversed) {
	br.cur = root
	br.stack = zhan.NewStack()
	return br
}
func (br *BSTIteratorReversed) hasPrev() bool {
	if br.cur != nil || !br.stack.IsEmpty() {
		return true
	} else {
		return false
	}
}
func (br *BSTIteratorReversed) prev() int {
	for br.cur != nil {
		br.stack.Push(br.cur)
		br.cur = br.cur.Right
	}
	node, _ := br.stack.Pop().(*shu.TreeNode)
	val := node.Val
	br.cur = node.Left
	return val
}
func findTarget2(root *shu.TreeNode, k int) bool {
	if root == nil {
		return false
	}
	iterNext := Constructor1(root)
	iterPrev := constructorBR(root)
	next := iterNext.next()
	prev := iterPrev.prev()
	for next != prev {
		if next+prev == k {
			return true
		}
		if next+prev < k {
			next = iterNext.next()
		}
		if next+prev > k {
			prev = iterPrev.prev()
		}
	}
	return false
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
	v1 := findTarget(tree, 12)
	v2 := findTarget(tree, 22)
	fmt.Println(v1, v2)
	v3 := findTarget2(tree, 12)
	v4 := findTarget2(tree, 22)
	fmt.Println(v3, v4)
}
