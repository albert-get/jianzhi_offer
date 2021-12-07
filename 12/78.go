package main

import (
	"container/heap"
	"fmt"
	"harry.com/jianzhi_offer/lianbiao"
)

/**
合并排序链表
*/
/**
题目：输入k
个排序的链表，请将它们合并成一个排序的链表。
*/

type heapNode []*lianbiao.ListNode

func (h heapNode) Len() int { return len(h) }

func (h heapNode) Less(i, j int) bool { return h[i].Val < h[j].Val }

func (h heapNode) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *heapNode) Pop() interface{} {
	old := *h
	n := len(old)

	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
func (h *heapNode) Peek() *lianbiao.ListNode {
	old := *h
	x := old[0]
	return x
}

func (h *heapNode) Push(x interface{}) {
	*h = append(*h, x.(*lianbiao.ListNode))
}

func mergeKLists(lists []*lianbiao.ListNode) *lianbiao.ListNode {
	dummy := new(lianbiao.ListNode)
	cur := dummy
	minHeap := heapNode{}
	heap.Init(&minHeap)
	for _, list := range lists {
		if list != nil {
			heap.Push(&minHeap, list)
		}
	}
	for minHeap.Len() != 0 {
		least, _ := heap.Pop(&minHeap).(*lianbiao.ListNode)
		cur.Next = least
		cur = least
		if least.Next != nil {
			heap.Push(&minHeap, least.Next)
		}
	}
	return dummy.Next
}
func mergeKLists2(lists []*lianbiao.ListNode) *lianbiao.ListNode {
	if len(lists) == 0 {
		return nil
	}
	return mergeList(lists, 0, len(lists))
}
func mergeList(lists []*lianbiao.ListNode, start int, end int) *lianbiao.ListNode {
	if start+1 == end {
		return lists[start]
	}
	mid := (start + end) / 2
	head1 := mergeList(lists, start, mid)
	head2 := mergeList(lists, mid, end)
	return mergel2(head1, head2)
}
func mergel2(head1 *lianbiao.ListNode, head2 *lianbiao.ListNode) *lianbiao.ListNode {
	dummy := new(lianbiao.ListNode)
	cur := dummy
	for head1 != nil && head2 != nil {
		if head1.Val < head2.Val {
			cur.Next = head1
			head1 = head1.Next
		} else {
			cur.Next = head2
			head2 = head2.Next
		}
		cur = cur.Next
	}
	if head1 == nil {
		cur.Next = head2
	} else {
		cur.Next = head1
	}
	return dummy.Next
}
func main() {
	l1 := lianbiao.Constructor([]int{1, 4, 7})
	l2 := lianbiao.Constructor([]int{2, 5, 8})
	l3 := lianbiao.Constructor([]int{3, 6, 9})
	lists := []*lianbiao.ListNode{l1, l2, l3}
	//r := mergeKLists(lists)
	r := mergeKLists2(lists)
	fmt.Println(r)
}
