package main

import (
	"fmt"
	"harry.com/jianzhi_offer/dui"
)

/**
朋友圈
*/
/**
题目：假设一个班级中有n
个学生。学生之间有些是朋友，有些不是。朋友关系是可以传递的。例如，A是B的直接朋友，B是C的直接朋友，那么A是C的间接朋友。定义朋友圈就是一组直接朋友或间接朋友的学生。输入一个n
×n
的矩阵M
表示班上的朋友关系，如果M
[i
][j
]=1，那么学生i
和学生j
是直接朋友。请计算该班级中朋友圈的数目。
*/

func findCircleNum(M [][]int) int {
	visited := make([]bool, len(M))
	result := 0
	for i := 0; i < len(M); i++ {
		if !visited[i] {
			findCircle(M, &visited, i)
			result++
		}
	}
	return result
}
func findCircle(M [][]int, visited *[]bool, i int) {
	queue := dui.QueueInt{}
	queue.Push(i)
	t := *visited
	t[i] = true
	for !queue.IsEmpty() {
		p := queue.Pop()
		for friend := 0; friend < len(M); friend++ {
			if M[p][friend] == 1 && !t[friend] {
				queue.Push(friend)
				t[friend] = true
			}
		}
	}
}
func findCircleNum1(M [][]int) int {
	fathers := make([]int, len(M))
	for i := 0; i < len(fathers); i++ {
		fathers[i] = i
	}
	count := len(M)
	for i := 0; i < len(M); i++ {
		for j := i + 1; j < len(M); j++ {
			if M[i][j] == 1 && union(&fathers, i, j) {
				count--
			}
		}
	}
	return count
}
func findFather(fathers *[]int, i int) int {
	t := *fathers
	if t[i] != i {
		findFather(fathers, t[i])
	}
	return t[i]
}
func union(fathers *[]int, i int, j int) bool {
	t := *fathers
	fatherOfI := findFather(fathers, i)
	fatherOfJ := findFather(fathers, j)
	if fatherOfI != fatherOfJ {
		t[fatherOfI] = fatherOfJ
		return true
	}
	return false
}
func main() {
	//f:=findCircleNum([][]int{{1,1,0},{1,1,0},{0,0,1}})
	f := findCircleNum1([][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}})
	fmt.Println(f)
}
