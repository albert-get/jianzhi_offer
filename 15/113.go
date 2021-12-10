package main

import (
	"fmt"
	"harry.com/jianzhi_offer/dui"
)

/**
课程顺序
*/
/**
题目：n
门课程的编号为0～n
-1。输入一个数组prerequisites，它的每个元素prerequisites[i
]表示两门课程的先修顺序。如果prerequisites[i
]=[ai
，bi
]，那么必须先修完bi
才能修ai
。请根据总课程数n
和表示先修顺序的prerequisites得出一个可行的修课序列。如果有多个可行的修课序列，则输出任意一个可行的序列；如果没有可行的修课序列，则输出空序列。
*/

func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make(map[int][]int)
	for i := 0; i < numCourses; i++ {
		graph[i] = []int{}
	}
	inDegrees := make([]int, numCourses)
	for _, prereq := range prerequisites {
		graph[prereq[1]] = append(graph[prereq[1]], prereq[0])
		inDegrees[prereq[0]]++
	}
	queue := dui.QueueInt{}
	for i := 0; i < numCourses; i++ {
		if inDegrees[i] == 0 {
			queue.Push(i)
		}
	}
	var order []int
	for !queue.IsEmpty() {
		course := queue.Pop()
		order = append(order, course)
		for _, next := range graph[course] {
			inDegrees[next]--
			if inDegrees[next] == 0 {
				queue.Push(next)
			}
		}
	}
	if len(order) == numCourses {
		return order
	} else {
		return []int{0}
	}
}
func main() {
	r := findOrder(4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}})
	fmt.Println(r)
}
