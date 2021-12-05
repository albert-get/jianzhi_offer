package main

import (
	"container/heap"
	"debug/dwarf"
	"fmt"
	"harry.com/jianzhi_offer/duu"
	"math"
)

/**
出现频率最高的k个数字
*/
/**
题目：请找出数组中出现频率最高的k
个数字。例如，当k
等于2时，输入数组[1，2，2，1，3，1]，由于数字1出现了3次，数字2出现了2次，数字3出现了1次，因此出现频率最高的2个数字是1和2。
*/

func topKFrequent(nums []int, k int) []int {
	numToCount := make(map[int]int)
	for _, v := range nums {
		numToCount[v] += 1
	}
	minHeap := duu.MinHeapMap{}
	heap.Init(&minHeap)
	for k, v := range numToCount {
		entry := dwarf.Field{Attr: dwarf.Attr(k), Val: v}
		if minHeap.Len() < k {
			heap.Push(&minHeap, entry)
		} else {
			if entry.Val.(int) > minHeap.Peek().Val.(int) {
				heap.Pop(&minHeap)
				heap.Push(&minHeap, entry)
			}
		}
	}
	var result []int
	for _, v := range minHeap {
		result = append(result, int(v.Attr))
	}
	return result
}

func kSmallestPairs(sum1 []int, sum2 []int, k int) [][]int {
	minHeap := duu.ConMinHeapArr([][]int{}, sum1, sum2)
	if len(sum2) > 0 {
		for i := 0; i < int(math.Min(float64(k), float64(len(sum1)))); i++ {
			heap.Push(&minHeap, []int{i, 0})
		}
	}
	var result [][]int
	for ; k > 0 && minHeap.Len() > 0; k-- {
		ids, _ := heap.Pop(&minHeap).([]int)
		result = append(result, []int{sum1[ids[0]], sum2[ids[1]]})
		if ids[1] < len(sum2)-1 {
			heap.Push(&minHeap, []int{ids[0], ids[1] + 1})
		}
	}
	return result
}
func main() {
	//r := topKFrequent([]int{1, 2, 2, 1, 3, 1}, 2)
	//fmt.Println(r)

	num1 := []int{1, 5, 13, 21}
	num2 := []int{2, 4, 9, 15}
	r := kSmallestPairs(num1, num2, 3)
	fmt.Println(r)
}
