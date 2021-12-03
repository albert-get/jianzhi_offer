package main

import (
	"fmt"
	"harry.com/jianzhi_offer/shu"
	"strconv"
	"strings"
)

/**
序列化和反序列化二叉树
*/
/**
题目：请设计一个算法将二叉树序列化成一个字符串，并能将该字符串反序列化出原来二叉树的算法。
*/

func serialize(root *shu.TreeNode) string {
	if root == nil {
		return "#"
	}
	leftStr := serialize(root.Left)
	rightStr := serialize(root.Right)
	return strconv.Itoa(root.Val) + "," + leftStr + "," + rightStr
}
func deserialize(data string) *shu.TreeNode {
	nodeStrs := strings.Split(data, ",")
	i := []int{0}
	return dfs(nodeStrs, i)
}
func dfs(strs []string, i []int) *shu.TreeNode {
	str := strs[i[0]]
	i[0]++
	if str == "#" {
		return nil
	}
	v, _ := strconv.Atoi(str)
	node := &shu.TreeNode{Val: v}
	node.Left = dfs(strs, i)
	node.Right = dfs(strs, i)
	return node
}
func main() {
	tree := shu.ConstructorC2(
		[]shu.ToTree{
			{Val: 6, Level: 3},
			{Val: 6, Level: 2},
			{Val: 6, Level: 3},
			{Val: 6, Level: 1},
			{Val: 6, Level: 2},
		}, 1)
	str := serialize(tree)
	fmt.Println(str)
	toTree := deserialize(str)
	fmt.Println(toTree)
}
