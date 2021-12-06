package main

import (
	"fmt"
	"math"
	"sort"
)

/**
合并区间
*/
/**
题目：输入一个区间的集合，请将重叠的区间合并。每个区间用两个数字比较，分别表示区间的起始位置和结束位置。例如，输入区间[[1，3]，[4，5]，[8，10]，[2，6]，[9，12]，[15，18]]，合并重叠的区间之后得到[[1，6]，[8，12]，[15，18]]。
*/

func merge(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] <= intervals[j][0]
	})
	var merged [][]int
	i := 0
	for i < len(intervals) {
		temp := []int{intervals[i][0], intervals[i][1]}
		j := i + 1
		for j < len(intervals) && intervals[j][0] <= temp[1] {
			temp[1] = int(math.Max(float64(temp[1]), float64(intervals[j][1])))
			j++
		}
		merged = append(merged, temp)
		i = j
	}
	return merged
}
func main() {
	m := merge([][]int{{1, 3}, {4, 5}, {8, 10}, {2, 6}, {9, 12}, {15, 18}})
	fmt.Println(m)
}
