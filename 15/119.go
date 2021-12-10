package main

import (
	"fmt"
	"harry.com/jianzhi_offer/dui"
	"harry.com/jianzhi_offer/set"
	"math"
)

/**
最长连续序列
*/
/**
题目：输入一个无序的整数数组，请计算最长的连续数值序列的长度。例如，输入数组[10，5，9，2，4，3]，则最长的连续数值序列是[2，3，4，5]，因此输出4。
*/

func longestConsecutive(nums []int) int {
	myset := set.NewSetInt()
	for _, num := range nums {
		myset.Add(num)
	}
	longest := 0
	for myset.Size() > 0 {
		longest = int(math.Max(float64(longest), float64(bfs(myset, myset.Next()))))
	}
	return longest
}
func bfs(myset *set.ISet, num int) int {
	queue := dui.QueueInt{}
	queue.Push(num)
	myset.Remove(num)
	length := 1
	for !queue.IsEmpty() {
		i := queue.Pop()
		neighbors := []int{i - 1, i + 1}
		for _, neighbor := range neighbors {
			if myset.Contains(neighbor) {
				queue.Push(neighbor)
				myset.Remove(neighbor)
				length++
			}
		}
	}
	return length
}
func longestConsecutive1(nums []int) int {
	fathers := make(map[int]int)
	counts := make(map[int]int)
	all := set.NewSetInt()
	for _, num := range nums {
		fathers[num] = num
		counts[num] = 1
		all.Add(num)
	}
	for _, num := range nums {
		if all.Contains(num + 1) {
			union3(&fathers, &counts, num, num+1)
		}
		if all.Contains(num - 1) {
			union3(&fathers, &counts, num, num-1)
		}
	}
	longest := 0
	for _, length := range counts {
		longest = int(math.Max(float64(longest), float64(length)))
	}
	return longest
}
func findFather3(fathers *map[int]int, i int) int {
	t := *fathers
	if t[i] != i {
		t[i] = findFather3(fathers, t[i])
	}
	return t[i]
}
func union3(fathers *map[int]int, counts *map[int]int, i int, j int) {
	t := *fathers
	fatherOfI := findFather3(fathers, i)
	fatherOfJ := findFather3(fathers, j)
	if fatherOfI != fatherOfJ {
		t[fatherOfI] = fatherOfJ
		r := *counts
		countOfI := r[fatherOfI]
		countOfJ := r[fatherOfJ]
		r[fatherOfJ] = countOfI + countOfJ
	}
}
func main() {
	//f := longestConsecutive([]int{10, 5, 9, 2, 4, 3})
	f := longestConsecutive1([]int{10, 5, 9, 2, 4, 3})
	fmt.Println(f)
}
