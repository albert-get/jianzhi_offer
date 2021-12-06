package main

import (
	"fmt"
	"harry.com/jianzhi_offer/lianbiao"
	"math"
)

/**
归并
*/
/**
链表排序
*/
/**
题目：输入一个链表的头节点，请将该链表排序。
*/
func sortArray1(nums []int) []int {
	length := len(nums)
	src := nums
	dst := make([]int, length)
	for seg := 1; seg < length; seg += seg {
		for start := 0; start < length; start += seg * 2 {
			mid := int(math.Min(float64(start+seg), float64(length)))
			end := int(math.Min(float64(start+seg*2), float64(length)))
			i := start
			j := mid
			k := start
			for i < mid || j < end {
				if j == end || (i < mid && src[i] < src[j]) {
					dst[k] = src[i]
					k++
					i++
				} else {
					dst[k] = src[j]
					k++
					j++
				}
			}
		}
		temp := src
		src = dst
		dst = temp
	}
	return src
}

func sortArray2(nums []int) []int {
	dst := make([]int, len(nums))
	copy(dst, nums)
	mergesort(nums, dst, 0, len(nums))
	return dst
}
func mergesort(src []int, dst []int, start int, end int) {
	if start+1 >= end {
		return
	}
	mid := (start + end) / 2
	mergesort(dst, src, start, mid)
	mergesort(dst, src, mid, end)
	i := start
	j := mid
	k := start
	for i < mid || j < end {
		if j == end || (i < mid && src[i] < src[j]) {
			dst[k] = src[i]
			k++
			i++
		} else {
			dst[k] = src[j]
			k++
			j++
		}
	}
}

func sortList(head *lianbiao.ListNode) *lianbiao.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	head1 := head
	head2 := split(head)

	head1 = sortList(head1)
	head2 = sortList(head2)

	return mergel(head1, head2)
}
func split(head *lianbiao.ListNode) *lianbiao.ListNode {
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	second := slow.Next
	slow.Next = nil
	return second
}
func mergel(head1 *lianbiao.ListNode, head2 *lianbiao.ListNode) *lianbiao.ListNode {
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
	//r := sortArray1([]int{4, 1, 5, 6, 2, 7, 8, 3})
	//r := sortArray2([]int{4, 1, 5, 6, 2, 7, 8, 3})
	//fmt.Println(r)

	head := lianbiao.Constructor([]int{3, 5, 1, 4, 2, 6})
	re := sortList(head)
	fmt.Println(head, re)
}
