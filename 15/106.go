package main

import (
	"fmt"
	"harry.com/jianzhi_offer/dui"
)

/**
二分图
*/
/**
题目：如果能将一个图中的节点分成A、B两个部分，使任意一条边的一个节点属于A而另一个节点属于B，那么该图就是一个二分图。输入一个由数组graph表示的图，graph[i
]中包含所有和节点i
相邻的节点，请判断该图是否为二分图。
*/

func isBipartite(graph [][]int) bool {
	size := len(graph)
	colors := make([]int, size)
	for i, _ := range colors {
		colors[i] = -1
	}
	for i := 0; i < size; i++ {
		if colors[i] == -1 {
			if !setColor(graph, colors, i, 0) {
				return false
			}
			//if !setColor1(graph, colors, i, 0) {
			//	return false
			//}
		}
	}
	return true
}
func setColor(graph [][]int, colors []int, i int, color int) bool {
	queue := dui.Queue{}
	queue.Push(i)
	colors[i] = color
	for queue.Len() != 0 {
		v, _ := queue.Pop().(int)
		for _, neighbor := range graph[v] {
			if colors[neighbor] >= 0 {
				if colors[neighbor] == colors[v] {
					return false
				}
			} else {
				queue.Push(neighbor)
				colors[neighbor] = 1 - colors[v]
			}
		}
	}
	return true
}
func setColor1(graph [][]int, colors []int, i int, color int) bool {
	if colors[i] >= 0 {
		return colors[i] == color
	}
	colors[i] = color
	for _, neighbor := range graph[i] {
		if !setColor1(graph, colors, neighbor, 1-color) {
			return false
		}
	}
	return true
}
func main() {
	f := isBipartite([][]int{{1, 3}, {0, 2}, {1, 3}, {0, 2}})
	//f:=isBipartite([][]int{{1,2,3},{0,2},{0,1,3},{0,2}})
	fmt.Println(f)
}
