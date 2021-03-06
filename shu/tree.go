package shu

import "harry.com/jianzhi_offer/zhan"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type ToTree struct {
	Level int
	Val   int
}

func ConstructorC2(list []ToTree, level int) *TreeNode {
	var head *TreeNode
	for i, v := range list {
		if v.Level == level {
			head = &TreeNode{Val: v.Val}
			leftList := list[0:i]
			if len(leftList) > 0 {
				head.Left = ConstructorC2(leftList, level+1)
			}
			rightList := list[i+1 : len(list)]
			if len(rightList) > 0 {
				head.Right = ConstructorC2(rightList, level+1)
			}
			break
		}
	}
	return head
}

func InorderTraversalD(root *TreeNode) []int {
	nodes := &[]int{}
	dfs(root, nodes)
	return *nodes
}
func dfs(root *TreeNode, nodes *[]int) {
	if root != nil {
		dfs(root.Left, nodes)
		*nodes = append(*nodes, root.Val)
		dfs(root.Right, nodes)
	}
}

func InorderTraversalT(root *TreeNode) []int {
	var nodes []int
	stack := zhan.NewStack()
	cur := root
	for cur != nil || !stack.IsEmpty() {
		for cur != nil {
			stack.Push(cur)
			cur = cur.Left
		}
		left, _ := stack.Pop().(*TreeNode)
		nodes = append(nodes, left.Val)
		cur = left.Right
	}
	return nodes
}

func PreorderTraversalD(root *TreeNode) []int {
	nodes := &[]int{}
	dfsQ(root, nodes)
	return *nodes
}
func dfsQ(root *TreeNode, nodes *[]int) {
	if root != nil {
		*nodes = append(*nodes, root.Val)
		dfsQ(root.Left, nodes)
		dfsQ(root.Right, nodes)
	}
}
func PreorderTraversalT(root *TreeNode) []int {
	var result []int
	stack := zhan.NewStack()
	cur := root
	for cur != nil || !stack.IsEmpty() {
		for cur != nil {
			result = append(result, cur.Val)
			stack.Push(cur)
			cur = cur.Left
		}
		node := stack.Pop().(*TreeNode)
		cur = node.Right
	}
	return result
}

func PostorderTraversalD(root *TreeNode) []int {
	nodes := &[]int{}
	dfsP(root, nodes)
	return *nodes
}
func dfsP(root *TreeNode, nodes *[]int) {
	if root != nil {
		dfsP(root.Left, nodes)
		dfsP(root.Right, nodes)
		*nodes = append(*nodes, root.Val)
	}
}

func PostorderTraversalT(root *TreeNode) []int {
	var result []int
	stack := zhan.NewStack()
	cur := root
	var prev *TreeNode
	for cur != nil || !stack.IsEmpty() {
		for cur != nil {
			stack.Push(cur)
			cur = cur.Left
		}
		node, _ := stack.Peek().(*TreeNode)
		if node.Right != nil && node.Right != prev {
			cur = node.Right
		} else {
			stack.Pop()
			result = append(result, node.Val)
			prev = node
			cur = nil
		}
	}
	return result
}
type BSTIterator struct {
	cur   *TreeNode
	stack *zhan.Stack
}

func Constructor(root *TreeNode) BSTIterator {
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
	node, _ := bs.stack.Pop().(*TreeNode)
	val := node.Val
	bs.cur = node.Right
	return val
}

type BSTIteratorReversed struct {
	cur   *TreeNode
	stack *zhan.Stack
}

func constructorBR(root *TreeNode) (br BSTIteratorReversed) {
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
	node, _ := br.stack.Pop().(*TreeNode)
	val := node.Val
	br.cur = node.Left
	return val
}
