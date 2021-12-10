package main

import (
	"fmt"
	"math"
)

/**
多余的边
*/
/**
题目：树可以看成无环的无向图。在一个包含n
个节点（节点标号为从1到n
）的树中添加一条边连接任意两个节点，这棵树就会变成一个有环的图。给定一个在树中添加了一条边的图，请找出这条多余的边（用这条边连接的两个节点表示）。输入的图用一个二维数组edges表示，数组中的每个元素是一条边的两个节点[u
，v
]（u
＜v
）。如果有多个答案，请输出在数组edges中最后出现的边。
*/

func findRedundantConnection(edges [][]int) []int {
	maxVertex := 0
	for _, edge := range edges {
		maxVertex = int(math.Max(float64(maxVertex), float64(edge[0])))
		maxVertex = int(math.Max(float64(maxVertex), float64(edge[1])))
	}
	fathers := make([]int, maxVertex+1)
	for i := 1; i <= maxVertex; i++ {
		fathers[i] = i
	}
	for _, edge := range edges {
		if !Union2(&fathers, edge[0], edge[1]) {
			return edge
		}
	}
	return make([]int, 2)
}
func findFather2(fathers *[]int, i int) int {
	t := *fathers
	if t[i] != i {
		findFather2(fathers, t[i])
	}
	return t[i]
}
func Union2(fathers *[]int, i int, j int) bool {
	t := *fathers
	fatherOfI := findFather2(fathers, i)
	fatherOfJ := findFather2(fathers, j)
	if fatherOfI != fatherOfJ {
		t[fatherOfI] = fatherOfJ
		return true
	}
	return false
}
func main() {
	r := findRedundantConnection([][]int{{1, 2}, {1, 3}, {2, 4}, {3, 4}, {2, 5}})
	fmt.Println(r)
}
