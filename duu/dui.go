package duu

import (
	"container/heap"
	"debug/dwarf"
)

type IntHeap []int

func (h IntHeap) Len() int { return len(h) }

func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)

	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *IntHeap) Peek() int {
	old := *h
	x := old[0]
	return x
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

type IntHeapMax []int

func (h IntHeapMax) Len() int { return len(h) }

func (h IntHeapMax) Less(i, j int) bool { return h[i] > h[j] }

func (h IntHeapMax) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeapMax) Pop() interface{} {
	old := *h
	n := len(old)

	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *IntHeapMax) Peek() int {
	old := *h
	x := old[0]
	return x
}

func (h *IntHeapMax) Push(x interface{}) {
	*h = append(*h, x.(int))
}

type MinHeapMap []dwarf.Field

func (h MinHeapMap) Len() int { return len(h) }

func (h MinHeapMap) Less(i, j int) bool { return h[i].Val.(int) < h[j].Val.(int) }

func (h MinHeapMap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MinHeapMap) Pop() interface{} {
	old := *h
	n := len(old)

	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *MinHeapMap) Peek() dwarf.Field {
	old := *h
	x := old[0]
	return x
}

func (h *MinHeapMap) Push(x interface{}) {
	*h = append(*h, x.(dwarf.Field))
}

type MinHeapArr [][]int
func ConMinHeapArr(h MinHeapArr,nums1 []int,nums2 []int) MinHeapWA {
	m:=MinHeapWA{
		h:&h,
		nums1: nums1,
		nums2: nums2,
	}
	heap.Init(&m)
	return m
}
type MinHeapWA struct {
	h *MinHeapArr
	nums1 []int
	nums2 []int
}

func (h *MinHeapWA) Len() int { return len(*h.h) }

func (h *MinHeapWA) Less(i, j int) bool {
	hhh:=*h.h
	return h.nums1[hhh[i][0]] + h.nums2[hhh[i][1]] < h.nums1[hhh[j][0]] + h.nums2[hhh[j][1]]
}

func (h *MinHeapWA) Swap(i, j int) {
	hhh:=*h.h
	hhh[i], hhh[j] = hhh[j], hhh[i]
}

func (h *MinHeapWA) Pop() interface{} {
	old := *h.h
	n := len(old)

	x := old[n-1]
	*h.h = old[0 : n-1]
	return x
}
func (h *MinHeapWA) Peek() []int {
	old := *h.h
	x := old[0]
	return x
}

func (h *MinHeapWA) Push(x interface{}) {
	*h.h = append(*h.h, x.([]int))
}
