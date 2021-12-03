package shu

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
