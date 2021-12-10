package main

import (
	"fmt"
	"harry.com/jianzhi_offer/dui"
	"harry.com/jianzhi_offer/set"
	"reflect"
)

/**
重建序列
*/
/**
题目：长度为n
的数组org是数字1～n
的一个排列，seqs是若干序列，请判断数组org是否为可以由seqs重建的唯一序列。重建的序列是指seqs所有序列的最短公共超序列，即seqs中的任意序列都是该序列的子序列。
*/

func sequenceReconstruction(org []int, seqs [][]int) bool {
	graph := make(map[int]*set.ISet)
	inDegrees := make(map[int]int)
	for _, seq := range seqs {
		for _, num := range seq {
			if num < 1 || num > len(org) {
				return false
			}
			if _, ok := graph[num]; !ok {
				graph[num] = set.NewSetInt()
			}
			if _, ok := inDegrees[num]; !ok {
				inDegrees[num] = 0
			}
		}
		for i := 0; i < len(seq)-1; i++ {
			num1 := seq[i]
			num2 := seq[i+1]
			if !graph[num1].Contains(num2) {
				graph[num1].Add(num2)
				inDegrees[num2] += 1
			}

		}
	}
	queue := dui.QueueInt{}
	for num, _ := range inDegrees {
		if inDegrees[num] == 0 {
			queue.Push(num)
		}
	}
	var built []int
	for queue.Len() == 1 {
		num := queue.Pop()
		built = append(built, num)

		for next, _ := range graph[num].M {
			inDegrees[next] -= 1
			if inDegrees[next] == 0 {
				queue.Push(next)
			}
		}
	}
	return reflect.DeepEqual(built, org)
}
func main() {
	f := sequenceReconstruction([]int{4, 1, 5, 2, 6, 3}, [][]int{{5, 2, 6, 3}, {4, 1, 5, 2}})
	fmt.Println(f)
}
