package main

import "fmt"

/**
删除倒数第k个节点
*/
/**
题目：如果给定一个链表，请问如何删除链表中的倒数第k个节点？假设链表中节点的总数为n，那么1≤k≤n。要求只能遍历链表一次。
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	queue := &ListNode{Next: head}
	r := queue
	for i := 0; i < n; i++ {
		r = r.Next
	}
	l := queue
	for r.Next != nil {
		l = l.Next
		r = r.Next
	}
	l.Next = l.Next.Next
	return queue.Next
}
func Constructor(arr []int) *ListNode {
	head := ListNode{}
	p := &head
	for _, v := range arr {
		cur := &ListNode{Val: v}
		p.Next = cur
		p = p.Next
	}
	return head.Next
}
func main() {
	list := Constructor([]int{1, 2, 3, 4, 5})
	result := removeNthFromEnd(list, 2)
	for result != nil {
		fmt.Println(result.Val)
		result = result.Next
	}
}
