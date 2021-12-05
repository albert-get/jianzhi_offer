package main

import (
	"container/heap"
	"fmt"
	"harry.com/jianzhi_offer/duu"
)

/**
数据流的第k大数字
*/
/**
题目：请设计一个类型KthLargest，它每次从一个数据流中读取一个数字，并得出数据流已经读取的数字中第k（k≥1）大的数字。该类型的构造函数有两个参数：一个是整数k，另一个是包含数据流中最开始数字的整数数组nums（假设数组nums的长度大于k）。该类型还有一个函数add，用来添加数据流中的新数字并返回数据流中已经读取的数字的第k大数字。
*/

type KthLargest struct {
	minHeap duu.IntHeap
	size    int
}

func conKL(k int, nums []int) (kl KthLargest) {
	kl.size = k
	h := &duu.IntHeap{}
	heap.Init(h)
	kl.minHeap = *h
	for _, v := range nums {
		kl.add(v)
	}
	return kl
}
func (kl *KthLargest) add(val int) int {
	if kl.minHeap.Len() < kl.size {
		heap.Push(&kl.minHeap, val)
	} else if val > kl.minHeap.Peek() {
		heap.Pop(&kl.minHeap)
		heap.Push(&kl.minHeap, val)
	}
	return kl.minHeap.Peek()
}
func main() {
	h:=&duu.IntHeap{2, 1, 5, 6, 4, 3, 7, 9, 8, 0}
	heap.Init(h)
	v:=heap.Pop(h)
	fmt.Println(h,v)

	//kl := conKL(3, []int{4, 5, 8, 2})
	//v1 := kl.add(3)
	//v2 := kl.add(5)
	//fmt.Println(v1, v2)
}
