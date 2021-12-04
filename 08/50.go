package main

import (
	"fmt"
	"harry.com/jianzhi_offer/shu"
)

/**
向下的路径节点值之和
*/
/**
题目：给定一棵二叉树和一个值sum，求二叉树中节点值之和等于sum的路径的数目。路径的定义为二叉树中顺着指向子节点的指针向下移动所经过的节点，但不一定从根节点开始，也不一定到叶节点结束。
*/

func pathSum(root *shu.TreeNode, sum int) int {
	mymap := make(map[int]int)
	mymap[0] = 1
	return dfs2(root, sum, mymap, 0)
}
func dfs2(root *shu.TreeNode, sum int, myMap map[int]int, path int) int {
	if root == nil {
		return 0
	}
	path += root.Val
	count := myMap[path-sum]
	myMap[path] += 1
	count += dfs2(root.Left, sum, myMap, path)
	count += dfs2(root.Right, sum, myMap, path)
	myMap[path] -= 1
	return count
}

func main() {
	tree := shu.ConstructorC2(
		[]shu.ToTree{
			{Val: 1, Level: 3},
			{Val: 2, Level: 2},
			{Val: 6, Level: 3},
			{Val: 5, Level: 1},
			{Val: 3, Level: 3},
			{Val: 4, Level: 2},
			{Val: 7, Level: 3},
		}, 1)
	sum := pathSum(tree, 8)
	fmt.Println(sum)

}
