package main

import (
	"fmt"
	"harry.com/jianzhi_offer/dui"
	"harry.com/jianzhi_offer/shu"
)

type CBTInserter struct {
	queue dui.Queue
	root  *shu.TreeNode
}

func constructor(root *shu.TreeNode) (this CBTInserter) {
	this.root = root
	this.queue = dui.Queue{}
	this.queue.Push(root)
	head, _ := this.queue.Peek().(*shu.TreeNode)
	for head.Left != nil && head.Right != nil {
		node, _ := this.queue.Pop().(*shu.TreeNode)
		this.queue.Push(node.Left)
		this.queue.Push(node.Right)
		head, _ = this.queue.Peek().(*shu.TreeNode)
	}
	return this
}
func (cb *CBTInserter) insert(v int) int {
	parent, _ := cb.queue.Peek().(*shu.TreeNode)
	node := &shu.TreeNode{Val: v}
	if parent.Left == nil {
		parent.Left = node
	} else {
		parent.Right = node
		cb.queue.Pop()
		cb.queue.Push(parent.Left)
		cb.queue.Push(parent.Right)
	}
	return parent.Val
}

func main() {
	tree := shu.ConstructorC2(
		[]shu.ToTree{
			{Val: 4, Level: 3},
			{Val: 2, Level: 2},
			{Val: 5, Level: 3},
			{Val: 1, Level: 1},
			{Val: 6, Level: 3},
			{Val: 3, Level: 2},
		}, 1)
	fmt.Println(tree)
	cb := constructor(tree)
	cb.insert(7)
	cb.insert(8)
	cb.insert(9)
	fmt.Println(cb)
}
