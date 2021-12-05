package main

import (
	"container/heap"
	"fmt"
	"harry.com/jianzhi_offer/duu"
	"math"
)

/**
和最小的k个数对
*/
/**
题目：给定两个递增排序的整数数组，从两个数组中各取一个数字u
和v
组成一个数对（u
，v
），请找出和最小的k
个数对。例如，输入两个数组[1，5，13，21]和[2，4，9，15]，和最小的3个数对为（1，2）、（1，4）和（2，5）。
*/

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
func kSmallestPairs2(nums1 []int, nums2 []int, k int) [][]int {

	maxHeap := &duu.MinHeapArr{}
	for i := 0; i < int(math.Min(float64(k), float64(len(nums1)))); i++ {
		for j := 0; j < int(math.Min(float64(k), float64(len(nums2)))); j++ {
			if maxHeap.Len() >= k {
				root := maxHeap.Peek()
				if root[0]+root[1] > nums1[i]+nums2[i] {
					heap.Pop(maxHeap)
					heap.Push(maxHeap, []int{nums1[i], nums2[j]})
				}
			} else {
				heap.Push(maxHeap, []int{nums1[i], nums2[j]})
			}
		}
	}
	var result [][]int
	for maxHeap.Len() > 0 {
		vals, _ := heap.Pop(maxHeap).([]int)
		result = append(result, vals)
	}
	return result
}

func main() {
	num1 := []int{1, 5, 13, 21}
	num2 := []int{2, 4, 9, 15}
	r := kSmallestPairs(num1, num2, 3)
	r2 := kSmallestPairs2(num1, num2, 3)
	fmt.Println(r, r2)
}
